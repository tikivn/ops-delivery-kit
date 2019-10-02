package oms

type Order struct {
	Code            string          `json:"code"`
	Status          string          `json:"status"`
	Substatus       string          `json:"substatus"`
	IsRMA           bool            `json:"is_rma"`
	GrandTotal      float64         `json:"grand_total"`
	Subtotal        float64         `json:"subtotal"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
	ShippingPlan    ShippingPlan    `json:"shipping_plan"`
	Payment         Payment         `json:"payment"`
	Items           []Item          `json:"items"`
	Warehouse       Warehouse       `json:"warehouse"`
	Shipment        ShipmentWrap    `json:"shipment"`
	BackendID       int64           `json:"backend_id"`
	OrderID         int64           `json:"order_id"`

	// =====================
	// Unused fields & sensitive data, should not parse
	// =====================

	// AppliedRuleIDS      string          `json:"applied_rule_ids"`
	// CouponCode          string          `json:"coupon_code"`
	// CreatedAt           TimeWrap        `json:"created_at"`
	// Customer            Customer        `json:"customer"`
	// DeliveryConfirmed   int             `json:"delivery_confirmed"`
	// DeliveryConfirmedAt TimeWrap        `json:"delivery_confirmed_at"`
	// DeliveryNote        string          `json:"delivery_note"`
	// DiscountAmount      int64           `json:"discount_amount"`
	// DiscountCoupon      int64           `json:"discount_coupon"`
	// DiscountOther       int64           `json:"discount_other"`
	// DiscountTikixu      int64           `json:"discount_tikixu"`
	// Extra               json.RawMessage `json:"extra"`
	FulfillmentType string `json:"fulfillment_type"`
	// GiftCardCode        string          `json:"gift_card_code"`
	// GiftCardMount       int64           `json:"gift_card_mount"`
	// InventoryStatus     string          `json:"inventory_status"`
	// IsBookcare          string          `json:"is_bookcare"`
	// ItemsCount          int64           `json:"items_count"`
	// ItemsQty            int64           `json:"items_qty"`
	// LinkedCode          string          `json:"linked_code"`
	// OriginalCode        string          `json:"original_code"`
	// Platform            string          `json:"platform"`
	// PurchasedAt         TimeWrap        `json:"purchased_at"`
	// RelationCode        string          `json:"relation_code"`
	// State               string          `json:"state"`
	// Substate            string          `json:"substate"`
	// TaxInfo             TaxInfo         `json:"tax_info"`
	// Transactions        []Transaction   `json:"transactions"`
	// Type                string          `json:"type"`
	// UpdatedAt           TimeWrap        `json:"updated_at"`
}

// type Customer struct {
// 	CustomerID int64  `json:"customer_id"`
// 	FullName   string `json:"full_name"`
// 	IP         string `json:"ip"`
// 	Phone      string `json:"phone"`
// 	Email      string `json:"email"`
// 	GroupID    int64  `json:"group_id"`
// 	GroupName  string `json:"group_name"`
// }

type Item struct {
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	ProductSku  string `json:"product_sku"`
	Price       int64  `json:"price"`
	Qty         int64  `json:"qty"`
	ProductType string `json:"product_type"`
	// =====================
	// Unused fields & sensitive data, should not parse
	// =====================

	// AppliedRuleIDS   string          `json:"applied_rule_ids"`
	// BackendID        int64           `json:"backend_id"`
	// CatalogGroupName string          `json:"catalog_group_name"`
	// DiscountAmount   int64           `json:"discount_amount"`
	// DiscountData     string          `json:"discount_data"`
	// DiscountOther    int64           `json:"discount_other"`
	// DiscountTikixu   int64           `json:"discount_tikixu"`
	// Extra            json.RawMessage `json:"extra"`
	// FulfilledAt      TimeWrap   `json:"fulfilled_at"`
	// InventoryType    string          `json:"inventory_type"`
	// IsBookCare       bool            `json:"is_book_care"`
	// IsEbook          bool            `json:"is_ebook"`
	// IsFreeGift       bool            `json:"is_free_gift"`
	// IsFulfilled      bool            `json:"is_fulfilled"`
	// IsTaxable        bool            `json:"is_taxable"`
	// ItemID           int64           `json:"item_id"`
	// ParentItemID     int64           `json:"parent_item_id"`
	// ProductMasterID  int64           `json:"product_master_id"`
	// ProductSuperID   int64           `json:"product_super_id"`

	// RowTotal         int64           `json:"row_total"`
	// SellerID         int64           `json:"seller_id"`
	// SellerName       string          `json:"seller_name"`
	// Subtotal         int64           `json:"subtotal"`
}

func (i Item) IsVirtualProduct() bool {
	return i.ProductType == "virtual"
}

type Payment struct {
	Method    string `json:"method"`
	IsPrepaid bool   `json:"is_prepaid"`

	// =====================
	// Unused fields & sensitive data, should not parse
	// =====================

	// PaymentID      int64  `json:"payment_id"`
	// SelectedMethod string `json:"selected_method"`
	// Status         string `json:"status"`
	// Description    string `json:"description"`
	// BackendID      int64  `json:"backend_id"`
}

type Shipment struct {
	PartnerID    string `json:"partner_id"`
	PartnerName  string `json:"partner_name"`
	TrackingCode string `json:"tracking_code"`
	Status       string `json:"status"`
}

type ShippingAddress struct {
	FullName     string `json:"full_name"`
	Street       string `json:"street"`
	District     string `json:"district"`
	Ward         string `json:"ward"`
	WardTikiCode string `json:"ward_tiki_code"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`

	// =====================
	// Unused fields & sensitive data, should not parse
	// =====================

	// ShippingAddressID int64  `json:"shipping_address_id"`
	// WardID            int64  `json:"ward_id"`
	// DistrictID        int64  `json:"district_id"`
	// DistrictTikiCode  string `json:"district_tiki_code"`
	// RegionID          int64  `json:"region_id"`
	// RegionTikiCode    string `json:"region_tiki_code"`
	// BackendID         int64  `json:"backend_id"`
}

type ShippingPlan struct {
	PlanID               int      `json:"plan_id"`
	PlanName             string   `json:"plan_name"`
	PromisedDeliveryDate TimeWrap `json:"promised_delivery_date"`

	// =====================
	// Unused fields & sensitive data, should not parse
	// =====================

	// DeliveryCommitmentTime string `json:"delivery_commitment_time"`
	// IsFreeShipping         bool   `json:"is_free_shipping"`
	// ShippingAmount         int64  `json:"shipping_amount"`
	// HandlingFee            int64  `json:"handling_fee"`
	// Description            string `json:"description"`
}

// type TaxInfo struct {
// 	CompanyName string `json:"company_name"`
// 	TaxCode     string `json:"tax_code"`
// 	Address     string `json:"address"`
// }

// type Transaction struct {
// 	TransactionID        int64   `json:"transaction_id"`
// 	GatewayName          string  `json:"gateway_name"`
// 	TransactionIdentity  string  `json:"transaction_identity"`
// 	GatewayTransactionID string  `json:"gateway_transaction_id"`
// 	ReferenceNumber      string  `json:"reference_number"`
// 	State                string  `json:"state"`
// 	Description          string  `json:"description"`
// 	Details              Details `json:"details"`
// 	CreatedAt            string  `json:"created_at"`
// 	UpdatedAt            string  `json:"updated_at"`
// }

// type Details struct {
// 	CardNumber     string `json:"card_number"`
// 	CardNameHolder string `json:"card_name_holder"`
// }

type Warehouse struct {
	WarehouseID   int    `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
}
