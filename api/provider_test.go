package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
)

func TestProviderAPI(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	providerInfo := mytype.ProviderInfo{
		ClientID:     randstr.Hex(10),
		ClientSecret: randstr.Hex(20),
		ProviderName: "test",
	}

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(mytype.AddProviderRequest{
			ClientID:     providerInfo.ClientID,
			ClientSecret: providerInfo.ClientSecret,
			ProviderName: providerInfo.ProviderName,
		},
		).Post("http://localhost:50060/v1/add/provider")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		response := mytype.AddProviderResponse{}
		err := json.Unmarshal(resp.Body(), &response)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response.Info.ProviderID, uuid.UUID{})
			assert.Equal(t, response.Info.ClientID, providerInfo.ClientID)
			assert.Equal(t, response.Info.ClientSecret, providerInfo.ClientSecret)
			assert.Equal(t, response.Info.ProviderName, providerInfo.ProviderName)
			providerInfo.ProviderID = response.Info.ProviderID
		}
	}

	providerInfo.ProviderLogo = "test"
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(mytype.UpdateProviderRequest{
			Info: providerInfo,
		},
		).Post("http://localhost:50060/v1/update/provider")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		response := mytype.UpdateProviderResponse{}
		err := json.Unmarshal(resp1.Body(), &response)
		if assert.Nil(t, err) {
			assert.Equal(t, response.Info.ProviderID, providerInfo.ProviderID)
			assert.Equal(t, response.Info.ClientID, providerInfo.ClientID)
			assert.Equal(t, response.Info.ClientSecret, providerInfo.ClientSecret)
			assert.Equal(t, response.Info.ProviderName, providerInfo.ProviderName)
			assert.Equal(t, response.Info.ProviderLogo, providerInfo.ProviderLogo)
		}
	}

	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(mytype.GetProvidersRequest{}).Post("http://localhost:50060/v1/get/all/providers")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
		response := mytype.GetProvidersResponse{}
		err := json.Unmarshal(resp2.Body(), &response)
		if assert.Nil(t, err) {
			assert.Equal(t, response.Infos[0].ProviderID, providerInfo.ProviderID)
			assert.Equal(t, response.Infos[0].ClientID, providerInfo.ClientID)
			assert.Equal(t, response.Infos[0].ClientSecret, providerInfo.ClientSecret)
			assert.Equal(t, response.Infos[0].ProviderName, providerInfo.ProviderName)
			assert.Equal(t, response.Infos[0].ProviderLogo, providerInfo.ProviderLogo)
		}
	}

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(mytype.DeleteProviderRequest{
			ProviderID: providerInfo.ProviderID,
		},
		).Post("http://localhost:50060/v1/delete/provider")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
	}
}
