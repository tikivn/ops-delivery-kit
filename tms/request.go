package tms

import (
	"time"
)

type Request struct {
	ClientCode              string       `json:"client_code"`
	ClientOrderID           string       `json:"client_order_id"`
	ServiceType             string       `json:"service_type"`
	ServiceCode             string       `json:"service_code"`
	RequestedTrackingNumber string       `json:"requested_tracking_number"`
	PartnerCode             string       `json:"partner_code"`
	BusinessType            BusinessType `json:"business_type"`
	RefType                 string       `json:"ref_type"`
	Reference               Reference    `json:"reference"`
	From                    AddressInfo  `json:"from"`
	To                      AddressInfo  `json:"to"`
	ReturnInfo              *AddressInfo `json:"return_info"`
	ParcelJob               ParcelJob    `json:"parcel_job"`
	IsThermalBagRequired    bool         `json:"is_thermal_bag_required"`
}

type TrackingInfo struct {
	ServiceType             string       `json:"service_type"`
	ServiceCode             string       `json:"service_code"`
	RequestedTrackingNumber string       `json:"requested_tracking_number" validate:"required"`
	PartnerCode             string       `json:"partner_code"`
	BusinessType            BusinessType `json:"business_type"`
	RefType                 string       `json:"ref_type"`
	Comment                 string       `json:"comment"`
}

type Product struct {
	ProductID   int64      `json:"product_id"`
	ProductSKU  string     `json:"product_sku"`
	ProductName string     `json:"product_name"`
	Qty         int        `json:"qty"`
	Price       float64    `json:"price"`
	Weight      float64    `json:"weight"`
	Dimensions  Dimensions `json:"dimensions"`
}

type Reference struct {
	TikiParcelNumber        string    `json:"tiki_parcel_number"`
	TikiOrderNumber         string    `json:"tiki_order_number"`
	NumberOfParcelsPerOrder int       `json:"number_of_parcels_per_order"`
	ParcelInfo              string    `json:"parcel_info"`
	Products                []Product `json:"products"`
}

type AddressInfo struct {
	Name        string      `json:"name"`
	PhoneNumber string      `json:"phone_number"`
	Email       string      `json:"email"`
	Address     string      `json:"address"`
	Ward        string      `json:"ward"`
	District    string      `json:"district"`
	Province    string      `json:"province"`
	Country     string      `json:"country"`
	AddressType string      `json:"address_type"`
	TikiCode    string      `json:"tiki_code"`
	ID          string      `json:"id"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ParcelJob struct {
	AllowWeekendDelivery  bool       `json:"allow_weekend_delivery"`
	CashOnDelivery        float64    `json:"cash_on_delivery"`
	InsuredValue          float64    `json:"insured_value"`
	DeliveryInstructions  string     `json:"delivery_instructions"`
	Dimensions            Dimensions `json:"dimensions"`
	IsPickupRequired      bool       `json:"is_pickup_required"`
	PickupInfo            string     `json:"pickup_info"`
	IsInstallRequired     bool       `json:"is_install_required"`
	ParcelTransferredTime time.Time  `json:"parcel_transferred_time"`
	DeliveryTimeslot      *TimeSlot  `json:"delivery_timeslot"`
	DeliveryStartTime     time.Time  `json:"delivery_start_time"`
	DeliveryDeadline      time.Time  `json:"delivery_deadline"`
}

type TimeSlot struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Timezone  string    `json:"timezone"`
}

type Dimensions struct {
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	Length float64 `json:"length"`
	Note   string  `json:"note"`
}

type ExtraInfo struct {
	PickupTime     time.Time `json:"pickup_time"`
	DeliveryTime   time.Time `json:"delivery_time"`
	ReturnTime     time.Time `json:"return_time"`
	CurrentStation string    `json:"current_station"`
	PreviousStatus string    `json:"previous_status"`
	SubStatus      string    `json:"sub_status"`
	NumberPick     int       `json:"number_pick"`
	NumberDeliver  int       `json:"number_deliver"`
	NumberReturn   int       `json:"number_return"`
	IsCrossZone    int       `json:"is_cross_zone"`
}

type Driver struct {
	Name         string  `json:"name"`
	Phone        string  `json:"phone"`
	LicensePlate string  `json:"license_plate"`
	PhotoURL     string  `json:"photo_url"`
	CurrentLat   float64 `json:"current_lat"`
	CurrentLng   float64 `json:"current_lng"`
}

type CallbackPayload struct {
	PartnerCode             string     `json:"partner_code"`
	RequestedTrackingNumber string     `json:"requested_tracking_number"`
	ReferenceTrackingNumber string     `json:"reference_tracking_number"`
	Status                  Status     `json:"status"`
	ReasonCode              Reason     `json:"reason_code"`
	Comment                 string     `json:"comment"`
	URI                     []string   `json:"uri"`
	RescheduleDelivery      time.Time  `json:"reschedule_delivery"`
	UpdateTime              time.Time  `json:"update_time"`
	SendTime                time.Time  `json:"send_time"`
	CashOnDelivery          float64    `json:"cash_on_delivery"`
	ShippingFee             float64    `json:"shipping_fee"`
	CodCollectionFee        float64    `json:"code_collection_fee"`
	ServiceFee              float64    `json:"service_fee"`
	ReturnFee               float64    `json:"return_fee"`
	TotalFee                float64    `json:"total_fee"`
	Dimensions              Dimensions `json:"dimensions"`
	ServiceType             string     `json:"service_type"`
	ServiceCode             string     `json:"service_code"`
	TikiOrderNumber         string     `json:"tiki_order_number"`
	ExtraInfo               ExtraInfo  `json:"extra_info"`
	RawData                 string     `json:"raw_data"`
	Version                 string     `json:"version"`
	Driver                  Driver     `json:"driver"`
	TrackingURL             string     `json:"tracking_url"`
}
