package logpose

type ActiveDriversResult struct {
	DriverID    string `json:"driver_id"`
	ThemisID    string `json:"themis_id"`
	DriverName  string `json:"driver_name"`
	Coordinates struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"coordinates"`
	LastCheckin string `json:"last_checkin"`
}
