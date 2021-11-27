package logia

import "time"

type Driver struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Phone        string  `json:"phone"`
	LicensePlate string  `json:"license_plate"`
	PhotoURL     string  `json:"photo_url"`
	CurrentLat   float64 `json:"current_lat"`
	CurrentLng   float64 `json:"current_lng"`
}

type CallbackPayload struct {
	RequestedTrackingNumber string    `json:"requested_tracking_number"`
	Status                  string    `json:"status"`
	ReasonCode              string    `json:"reason_code"`
	Comment                 string    `json:"comment"`
	UpdateTime              time.Time `json:"update_time"`
	Driver                  Driver    `json:"driver"`
	TrackingURL             string    `json:"tracking_url"`
}
