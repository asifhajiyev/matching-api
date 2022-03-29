package util

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

func StringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Errorf("'%s' could not be parsed", s)
	}
	return f
}
