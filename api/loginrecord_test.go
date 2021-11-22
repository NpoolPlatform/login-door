package api

import (
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/login-door/pkg/location"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLoginRecordAPI(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userID := uuid.New().String()
	appID := uuid.New().String()
	resp, err := location.GetLocationByIP("218.77.129.195")
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}

	cli := resty.New()
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(mytype.GetUserLoginRecordsRequest{
			UserID: userID,
			AppID:  appID,
		}).Post("http://localhost:50060/v1/get/user/login/records")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
	}

	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(mytype.GetAppLoginRecordsRequest{
			AppID: appID,
		}).Post("http://localhost:50060/v1/get/app/login/records")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
	}

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(mytype.GetLoginRecordsRequest{}).
		Post("http://localhost:50060/v1/get/all/login/records")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
	}
}
