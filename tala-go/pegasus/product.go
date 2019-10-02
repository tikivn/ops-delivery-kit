package pegasus

type DeliveryAttribute struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Product struct {
	AllowSetup          int                 `json:"allow_setup"`
	DeliveryAttributes  []DeliveryAttribute `json:"delivery_attributes"`
	ProductsetGroupName string              `json:"productset_group_name"`
	ProductsetID        int                 `json:"productset_id"`
	ProductsetName      string              `json:"productset_name"`
	SKU                 string              `json:"sku"`
	Name                string              `json:"name"`

	CurrentPrice int64 `json:"current_price"`
	ListPrice    int64 `json:"list_price"`
	Price        int64 `json:"price"`
}
