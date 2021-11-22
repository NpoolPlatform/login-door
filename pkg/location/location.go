package location

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/xerrors"
)

type IPLocation struct {
	Status      string
	Country     string
	CountryCode string
	Region      string
	RegionName  string
	City        string
	Zip         string
	Lat         float32
	Lon         float32
	Timezone    string
	Isp         string
	Org         string
	As          string
	Query       string
}

func GetLocationByIP(ip string) (*IPLocation, error) {
	cli := resty.New()
	cli.SetTimeout(10 * time.Second)

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return &IPLocation{}, err
	}

	ipLocation := IPLocation{}
	err = json.Unmarshal(resp.Body(), &ipLocation)
	if err != nil {
		return &IPLocation{}, err
	}

	if ipLocation.Status != "success" {
		return &IPLocation{}, xerrors.Errorf("fail to get ip location")
	}

	return &ipLocation, nil
}
