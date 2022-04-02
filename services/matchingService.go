package services

import (
	"github.com/asifhajiyev/matching-api/clients"
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
	radius := radiusToSearchDriver
	searchDriver, err := request.NewSearchDriverRequest(longitude, latitude, radius)
	if err != nil {
		return nil, err
	}
	r, err := ms.Client.Search(*searchDriver)

	if err != nil {
		return nil, err
	}

	if r.Data == nil {
		return nil, &error.Error{
			Code:    r.Code,
			Message: r.Message,
		}
	}

	ri := response.RideInfo{}
	util.InterfaceToStruct(r.Data, &ri)
	dlr := response.SearchDriverResponse{RideInfo: ri}

	return &dlr, nil
}
