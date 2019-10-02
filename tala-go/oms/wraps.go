package oms

import (
	"encoding/json"
	"time"
)

const (
	mysqlLayout = "2006-01-02 15:04:05"
)

type TimeWrap struct {
	time.Time
}

func (u *TimeWrap) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	// Ignore null, like in the main JSON package.
	if str == "null" {
		return nil
	}

	if str == `""` || str == `"0000-00-00 00:00:00"` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	u.Time, err = time.ParseInLocation(`"`+mysqlLayout+`"`, str, time.Local)
	return err
}

type ShipmentWrap struct {
	Shipment
}

func (u *ShipmentWrap) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	if str == "null" {
		return nil
	}
	// PHP empty object
	if str == "[]" {
		return nil
	}
	return json.Unmarshal(data, &u.Shipment)
}
