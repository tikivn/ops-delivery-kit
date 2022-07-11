package tms

import (
	"fmt"
	"time"
)

const (
	STACODCollected       stAction = "cod_collected"
	STASuccessfulDelivery stAction = "successful_delivery"
)

// Shipment tracking action
type stAction string

// StTimestamp Shipment tracking timestamp
type StTimestamp struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in format yyyy-MM-dd HH:mm:ss
func (t StTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))), nil
}

type ShipmentTracking struct {
	RequestCode string      `json:"request_code"`
	RefCode     string      `json:"ref_code"`
	BoxCode     string      `json:"box_code"`
	Action      stAction    `json:"action"`
	Timestamp   StTimestamp `json:"timestamp"`
	PartnerID   string      `json:"partner_id"`
	DriverID    string      `json:"driver_id"`
	HubID       string      `json:"hub_id"`
	ClientName  string      `json:"client_name"`
	COD         float64     `json:"cod"`
	Status      string      `json:"status"`
	SubStatus   string      `json:"sub_status"`
	TaskType 	string		`json:"task_type"`
}
