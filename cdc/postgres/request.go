package postgres

type ChangedData struct {
	Payload ChangedDataPayload `json:"payload"`
}

type ChangedDataPayload struct {
	After map[string]string `json:"after"`
}
