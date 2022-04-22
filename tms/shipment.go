package tms

type ShipmentTracking struct {
	RefCode   string `json:"ref_code"`
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"` // yyyy-MM-dd HH:mm:ss
	PartnerID string `json:"partner_id"`
	DriverID  string `json:"driver_id"`
	HubID     string `json:"hub_id"`
}
