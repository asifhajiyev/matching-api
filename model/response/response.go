package response

type RestResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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
