package util

import "github.com/sirupsen/logrus"

func InArrayString(arr []string, item string) bool {
	if len(arr) == 0 {
		return false
	}

	for _, value := range arr {
		if value == item {
			return true
		}
	}

	return false
}

func LogErrWithFields(err error, msg string, fields logrus.Fields) {
	logrus.
		WithError(err).
		WithFields(fields).
		Errorf(msg)
}
