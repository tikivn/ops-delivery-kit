package logpose

type ActiveDriversResult struct {
	DriverID    string      `json:"driver_id"`
	ThemisID    string      `json:"themis_id"`
	DriverName  string      `json:"driver_name"`
	Coordinates Coordinates `json:"coordinates"`
	LastCheckin string      `json:"last_checkin"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Boundary struct {
	NorthEast []float64 `json:"northeast"`
	SouthWest []float64 `json:"southwest"`
}

type GeocodePayload struct {
	Boundary *Boundary `json:"boundary"`
	Province string    `json:"province"`
	District string    `json:"district"`
	Ward     string    `json:"ward"`
	Street   string    `json:"street"`
}
