package lyoko

type SmsPayload struct {
	OrderCode string `json:"order_code"`
	Content   string `json:"content"`
}
