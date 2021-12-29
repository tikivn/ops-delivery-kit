package postgres

type ChangedData struct {
	Payload ChangedDataPayload
}

type ChangedDataPayload struct {
	After map[string]string
}
