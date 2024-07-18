package utils

import "time"

func GetCurrentISTTime() (time.Time, error) {
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return time.Time{}, err
	}

	istTime := time.Now().In(istLocation)
	return istTime, nil
}