package tms

import (
	"time"

	"github.com/google/uuid"
)

type RequestRefType string

type BusinessType string

type CreateShipmentResult struct {
	Status        string `json:"status,omitempty"`
	Error         string `json:"error,omitempty"`
	ZoneCode      string `json:"zone_code,omitempty"`
	PartnerCode   string `json:"partner_code,omitempty"`
	SortingCode   string `json:"sorting_code,omitempty"`
	TrackingID    string `json:"tracking_id,omitempty"`
	CurrentStatus string `json:"current_status"`
	ReasonCode    string `json:"reason_code"`
	Quote         Quote  `json:"quote"`
	TrackingURL   string `json:"tracking_url"`
}

type QuotesResult struct {
	Error  string  `json:"error,omitempty"`
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	PartnerCode       string            `json:"partner_code"`
	Service           Service           `json:"service"`
	Fee               Fee               `json:"fee"`
	EstimatedTimeline EstimatedTimeline `json:"estimated_timeline"`
	Distance          int               `json:"distance"`
}

type Service struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type Fee struct {
	Amount   float64 `json:"amount"`
	UnitCode string  `json:"unit_code"`
}

type EstimatedTimeline struct {
	Pickup  time.Time `json:"pickup"`
	Dropoff time.Time `json:"dropoff"`
}

type InfoResult struct {
	Error         string `json:"error,omitempty"`
	Quote         Quote  `json:"quote"`
	CurrentStatus string `json:"current_status"`
	ReasonCode    string `json:"reason_code"`
	Driver        Driver `json:"driver"`
	TrackingURL   string `json:"tracking_url"`
}

type ActiveDriversResult struct {
	Status string `json:"status"`
	Error  struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
	Data struct {
		DriverIds []uuid.UUID `json:"driver_ids"`
	} `json:"data,omitempty"`
}
