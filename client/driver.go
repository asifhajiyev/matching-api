package client

import (
	"fmt"
	"github.com/asifhajiyev/matching-api/model"
	"github.com/go-resty/resty/v2"
)

type DriverSearcher interface {
	Search(sd model.SearchDriver) (*RideInfo, error)
}
type driverSearch struct {
	Client *resty.Client
}

func NewDriverClient(client *resty.Client) DriverSearcher {
	return driverSearch{Client: client}
}

type RideInfo struct {
	DriverInfo DriverInfo `json:"driverInfo"`
	Distance   float64    `json:"distance"`
}

type DriverInfo struct {
	Location Location `json:"location"`
}

type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func (ds driverSearch) Search(sd model.SearchDriver) (*RideInfo, error) {
	var ri *RideInfo

	response, err := ds.Client.R().SetBody(sd).SetResult(ri).
		Post("http://localhost:8080/api/driver-location/search")

	fmt.Println("in client response", response)
	fmt.Println("in client err", err)

	if err != nil {
		return nil, err
	}
	return ri, nil
}
