package util

import (
	"time"

	"github.com/sirupsen/logrus"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	logrus.Infof("%s took %s", name, elapsed)
}

func SubtractByEpochTime(epochTime int64, timestamp time.Time) time.Duration {
	timestampUTC := timestamp.UTC() // UTC + 0

	return time.Unix(epochTime, 0).Sub(timestampUTC)
}
