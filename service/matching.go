package service

import (
	"fmt"
	"github.com/asifhajiyev/matching-api/client"
	"github.com/asifhajiyev/matching-api/error"
	"github.com/asifhajiyev/matching-api/model"
	"github.com/asifhajiyev/matching-api/util"
)

type MatchingService interface {
	Match(longitude, latitude string) *error.Error
}

type matchingService struct {
	Client client.DriverSearcher
}

func NewMatchingService(client client.DriverSearcher) MatchingService {
	return matchingService{Client: client}
}

func (ms matchingService) Match(longitude, latitude string) *error.Error {
	lng := util.StringToFloat(longitude)
	lt := util.StringToFloat(latitude)
	radius := radiusToSearchDriver

	searchDriver := model.SearchDriver{
		Longitude: lng,
		Latitude:  lt,
		Radius:    radius,
	}
	response, err := ms.Client.Search(searchDriver)

	fmt.Println("in service response", response)
	fmt.Println("in service err", err)

	return nil
}
