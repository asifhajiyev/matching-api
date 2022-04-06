package clients

import (
	"github.com/asifhajiyev/matching-api/constants"
	err "github.com/asifhajiyev/matching-api/error"
	"github.com/asifhajiyev/matching-api/logger"
	"github.com/asifhajiyev/matching-api/model"
	"github.com/asifhajiyev/matching-api/model/request"
	"github.com/asifhajiyev/matching-api/util"
	"github.com/go-resty/resty/v2"
	"os"
)

type DriverSearcher interface {
	SearchDriver(sd request.SearchDriverRequest) (*model.RestResponse, *err.Error)
}

type driverSearch struct {
	Client *resty.Client
}

func NewDriverClient(client *resty.Client) DriverSearcher {
	return driverSearch{Client: client}
}

func (ds driverSearch) SearchDriver(sd request.SearchDriverRequest) (*model.RestResponse, *err.Error) {
	logger.Info("SearchDriver.begin")
	rr := model.RestResponse{}
	dlAUsername := os.Getenv("client_driver_location_api_username")
	dlAPassword := os.Getenv("client_driver_location_api_password")
	resp, e := ds.Client.R().SetBody(sd).SetBasicAuth(dlAUsername, dlAPassword).Post("drivers/search")

	if e != nil {
		return nil, err.ServerError(constants.ErrorDriverApiDoesNotRespond)
	}
	if e := util.JsonToStruct(resp.Body(), &rr); e != nil {
		return nil, e
	}
	logger.Info("SearchDriver.end", &rr)
	return &rr, nil
}
