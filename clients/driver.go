package clients

import (
	err "github.com/asifhajiyev/matching-api/error"
	"github.com/asifhajiyev/matching-api/model/request"
	"github.com/asifhajiyev/matching-api/model/response"
	"github.com/asifhajiyev/matching-api/util"
	"github.com/go-resty/resty/v2"
)

type DriverSearcher interface {
	Search(sd request.SearchDriverRequest) (*response.RestResponse, *err.Error)
}

type driverSearch struct {
	Client *resty.Client
}

func NewDriverClient(client *resty.Client) DriverSearcher {
	return driverSearch{Client: client}
}

func (ds driverSearch) Search(sd request.SearchDriverRequest) (*response.RestResponse, *err.Error) {
	rr := response.RestResponse{}
	resp, e := ds.Client.R().SetBody(sd).Post("drivers/search")

	if e != nil {
		return nil, err.ServerError(e.Error())
	}
	util.JsonToStruct(resp.Body(), &rr)

	return &rr, nil
}