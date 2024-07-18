package seq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/GriffyHome/go-skeleton/pkg/config"
	"github.com/GriffyHome/go-skeleton/pkg/constants"
	"github.com/GriffyHome/go-skeleton/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitSeqLogger() {
	hook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		_, filePath, line, _ := runtime.Caller(4)
		fileName := path.Base(filePath)
		folderName := path.Dir(filePath)

		timestamp, _ := utils.GetCurrentISTTime()
		sendToSeq(level, msg, folderName, fileName, line, timestamp)
	})
	log.Logger = log.Hook(hook)
}

func sendToSeq(level zerolog.Level, msg string, folder string, file string, line int, timestamp time.Time) {

	seqUrl := config.GetSeqURL()

	logEvent := map[string]interface{}{
		"Timestamp":       timestamp.Format(time.RFC3339Nano),
		"MessageTemplate": msg,
		"Level":           strings.ToUpper(level.String()),
		"Properties": map[string]interface{}{
			"Folder": folder,
			"File":   file + ":" + strconv.Itoa(line),
			"Server": config.GetServiceID(),
		},
	}

	payload := map[string]interface{}{
		"Events": []map[string]interface{}{logEvent},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Error().Msgf("Failed to marshal log payload: %s\n", err.Error())
		return
	}

	req, err := http.NewRequest(http.MethodPost, seqUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Error().Msgf("Failed to create HTTP request: %s\n", err.Error())
		return
	}
	req.Header.Set(constants.ContentType, constants.ApplicationJSON)
	req.Header.Set(constants.SeqApiKey, config.GetSeqApiKey())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("Failed to send log to SEQ: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		responseBody, _ := io.ReadAll(resp.Body)
		log.Error().Msgf("Failed to send log to SEQ, status: %d, response: %s", resp.StatusCode, string(responseBody))
	}
}
