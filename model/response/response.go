package response

type SearchDriverResponse struct {
	RideInfo RideInfo `json:"rideInfo"`
}

type RideInfo struct {
	DriverInfo DriverInfo `json:"driverInfo"`
	Distance   float64    `json:"distance"`
}

type DriverInfo struct {
	Location Location `json:"location"`
}

type Location struct {
	Coordinates []float64 `json:"coordinates"`
}
