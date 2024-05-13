package strings

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

func ToInt(str string) int {
	id, err := strconv.Atoi(str)
	if err != nil {
		log.Errorln("Failed to convert string to int", err)
		return 0
	}

	return id
}

func FloatToString(input float64) string {
	return strconv.FormatFloat(input, 'f', -1, 64)
}
