package request

import (
	"github.com/asifhajiyev/matching-api/constants"
	error "github.com/asifhajiyev/matching-api/error"
	"github.com/asifhajiyev/matching-api/util"
)

type SearchDriverRequest struct {
	Radius      int        `json:"radius"`
	Coordinates Coordinate `json:"coordinates"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func NewSearchDriverRequest(longitude, latitude string, radius int) (*SearchDriverRequest, *error.Error) {
	lng, lngErr := util.StringToFloat(longitude)
	lt, ltErr := util.StringToFloat(latitude)

	if lngErr != nil || ltErr != nil {
		return nil, error.ValidationError(constants.ErrorUnprocessableCoordinates)
	}

	if !isValidLongitude(lng) && !isValidLatitude(lt) {
		return nil, error.ValidationError(constants.ErrorInvalidCoordinates)
	}

	return &SearchDriverRequest{
		Radius: radius,
		Coordinates: Coordinate{
			Longitude: lng,
			Latitude:  lt,
		},
	}, nil

}

func isValidLongitude(longitude float64) bool {
	return longitude >= -180 && longitude <= 180
}

func isValidLatitude(latitude float64) bool {
	return latitude >= -90 && latitude <= 90
}
