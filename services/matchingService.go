package services

import (
	"github.com/asifhajiyev/matching-api/clients"
	"github.com/asifhajiyev/matching-api/constants"
	"github.com/asifhajiyev/matching-api/error"
	"github.com/asifhajiyev/matching-api/model/request"
	"github.com/asifhajiyev/matching-api/model/response"
	"github.com/asifhajiyev/matching-api/util"
)

type MatchingService interface {
	Match(longitude, latitude string) (*response.SearchDriverResponse, *error.Error)
}

type matchingService struct {
	Client clients.DriverSearcher
}

func NewMatchingService(client clients.DriverSearcher) MatchingService {
	return matchingService{Client: client}
}

func (ms matchingService) Match(longitude, latitude string) (*response.SearchDriverResponse, *error.Error) {

	if longitude == "" || latitude == "" {
		return nil, error.ValidationError(constants.ErrorUnprocessableCoordinates)
	}

	radius := constants.RadiusToSearchDriver
	searchDriver, err := request.NewSearchDriverRequest(longitude, latitude, radius)
	if err != nil {
		return nil, err
	}

	driverResponse, err := ms.Client.SearchDriver(*searchDriver)
	if err != nil {
		return nil, err
	}

	if driverResponse.Data == nil {
		return nil, &error.Error{
			Code:    driverResponse.Code,
			Message: driverResponse.Message,
			Details: driverResponse.ErrorDetails,
		}
	}

	rideInfo := response.RideInfo{}
	if err = util.InterfaceToStruct(driverResponse.Data, &rideInfo); err != nil {
		return nil, err
	}

	searchDriverResponse := response.SearchDriverResponse{RideInfo: rideInfo}
	return &searchDriverResponse, nil
}
