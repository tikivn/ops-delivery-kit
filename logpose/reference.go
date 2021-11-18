package logpose

type ActiveDriversResult struct {
	DriverID    string `json:"driver_id"'`
	DriverName  string `json:"driver_name"'`
	Coordinates struct {
		Lat float32 `json:"lat"`
		Lng float32 `json:"lng"`
	} `json:"coordinates"`
	LastCheckin string `json:"last_checkin"`
}
