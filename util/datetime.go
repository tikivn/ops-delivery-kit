package util

import (
	"time"

	"github.com/sirupsen/logrus"
)

type DateTime struct {
	time.Time
}

const (
	mysqlLayout = "2006-01-02 15:04:05"
)

func (u *DateTime) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	// Ignore null, like in the main JSON package.
	if str == "null" {
		return nil
	}

	if str == `""` {
		return nil
	}

	// Fractional seconds are handled implicitly by Parse.
	u.Time, err = time.ParseInLocation(`"`+mysqlLayout+`"`, str, time.Local)
	return err
}

func TimeToPtr(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

func PtrToTime(t *time.Time, defaultValue time.Time) time.Time {
	if t == nil {
		return defaultValue
	}
	return *t
}

func TimeToSecond(t time.Time) time.Time {
	timeStampString := t.Format("2006-01-02 15:04:05")

	layOut := "2006-01-02 15:04:05"

	timeStamp, err := time.Parse(layOut, timeStampString)
	if err != nil {
		return t
	}

	return timeStamp
}

func MakeupDate(source time.Time, destination time.Time) time.Time {
	year, month, day := source.Date()
	hour, minute, _ := destination.Clock()

	return time.Date(year, month, day, hour, minute, 0, 0, source.Location())
}

func MaxTime(x, y time.Time) time.Time {
	if x.After(y) { // nếu ngày x sau ngày y
		return x
	}
	return y
}

func MinTime(x, y time.Time) time.Time {
	if x.After(y) { // nếu ngày x sau ngày y
		return y
	}
	return x
}

func SubTime(x, y time.Time) float64 {
	return x.Sub(y).Hours()
}

func ParseStringsToTime(dateString string, timeString string, timeZone string, date_layout string, time_layout string) (time.Time, error) {
	if dateString == "" {
		dateString = time.Time{}.Format(date_layout)
	}
	if timeString == "" {
		timeString = time.Time{}.Format(time_layout)
	}

	d, err := time.Parse(date_layout, dateString)
	if err != nil {
		logrus.WithError(err).Infof("%+v wrong time format", dateString)
		return time.Time{}, err
	}

	t, err := time.Parse(time_layout, timeString)
	if err != nil {
		logrus.WithError(err).Infof("%+v wrong time format", timeString)
		return time.Time{}, err
	}

	zone, err := time.LoadLocation(timeZone)
	if err != nil {
		logrus.WithError(err).Infof("%+v wrong time format", timeZone)
		return time.Time{}, err
	}

	return time.Date(d.Year(), d.Month(), d.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), zone), nil
}
