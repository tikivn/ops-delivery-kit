package luffy

import "time"

type RequestRefType string

const (
	RefTypeSellerReturn   RequestRefType = "seller_return"
	RefTypeSalesOrder     RequestRefType = "sales_order"
	RefTypeCustomerReturn RequestRefType = "customer_return"
	RefTypeSupplierReturn RequestRefType = "supplier_return"
)

type BusinessType string

const (
	B2C BusinessType = "B2C"
	C2C BusinessType = "C2C"
)

type CreateShipmentResult struct {
	Status      string `json:"status,omitempty"`
	Error       string `json:"error,omitempty"`
	ZoneCode    string `json:"zone_code,omitempty"`
	PartnerCode string `json:"partner_code,omitempty"`
	SortingCode string `json:"sorting_code,omitempty"`
	TrackingID  string `json:"tracking_id,omitempty"`
	Quote       Quote  `json:"quote"`
}

const (
	NJV  = "NJV"
	BEST = "BEST"
	GHN  = "GHN"
	JNT  = "JNT"
	GRAB = "GRAB"
)

type QuotesResult struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	Service           Service           `json:"service"`
	Fee               Fee               `json:"fee"`
	EstimatedTimeline EstimatedTimeline `json:"estimated_timeline"`
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
	Quote         Quote  `json:"quote"`
	CurrentStatus string `json:"current_status"`
}
