package tms

type BizError interface {
	error
	Code() string
}

type GetQuotesFailError struct {
	Message string
}

func (c GetQuotesFailError) Code() string {
	return "FAILED_TO_GET_QUOTES"
}

func (c GetQuotesFailError) Error() string {
	return c.Message
}

type CreateShipmentFailError struct {
	Message string
}

func (c CreateShipmentFailError) Code() string {
	return "FAILED_TO_CREATE_SHIPMENT"
}

func (c CreateShipmentFailError) Error() string {
	return c.Message
}

type CancelShipmentFailError struct {
	Message string
}

func (c CancelShipmentFailError) Code() string {
	return "FAILED_TO_CANCEL_SHIPMENT"
}

func (c CancelShipmentFailError) Error() string {
	return c.Message
}
