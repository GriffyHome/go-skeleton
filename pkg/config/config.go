package config

import (
	"fmt"
	"github.com/GriffyHome/go-skeleton/pkg/constants"
	errorlogs "github.com/GriffyHome/go-skeleton/pkg/constants/log_constants/errorLogs"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType(constants.DefaultConfigurationType)
	config.SetConfigName(env)
	config.AddConfigPath(constants.DefaultConfigurationPath)
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf(errorlogs.ParsingError, err.Error()))
	}
}


func GetConfig() *viper.Viper {
	return config
}

func DBConfig() DB {
	return DB{
		URL:      config.GetString("db.url"),
		Username: config.GetString("db.username"),
		Password: config.GetString("db.password"),
		Database: config.GetString("db.name"),
	}
}

func CassandraConfig() Cassandra {
	return Cassandra{
		Hosts:    config.GetStringSlice("cassandra.hosts"),
		Keyspace: config.GetString("cassandra.keyspace"),
		Username: config.GetString("cassandra.username"),
		Password: config.GetString("cassandra.password"),
	}
}

func GetSeqApiKey() string {
	return config.GetString("seq.api_key")
}

func GetSeqURL() string {
	return config.GetString("seq.url")

}

func GetServiceID() string {
	return config.GetString("service_id")
}

func GetRedisConnectionDetails() (host, password string) {
	return config.GetString("redis.host"), config.GetString("redis.password")
}