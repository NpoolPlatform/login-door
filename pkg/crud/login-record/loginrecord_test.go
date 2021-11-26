package loginrecord

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/login-door/pkg/location"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	testinit "github.com/NpoolPlatform/login-door/pkg/test-init"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestLoginRecordCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	userID := uuid.New().String()
	appID := uuid.New().String()
	resp, err := location.GetLocationByIP("218.77.129.195")
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}

	resp1, err := Create(context.Background(), &mytype.LoginRecord{
		UserID:    userID,
		AppID:     appID,
		IP:        resp.Query,
		Lat:       float64(resp.Lat),
		Lon:       float64(resp.Lon),
		LoginTime: uint32(time.Now().Unix()),
		Location:  resp.Timezone + " " + resp.Country + " " + resp.City,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp1)
	}

	resp2, err := GetByUser(context.Background(), &mytype.GetUserLoginRecordsRequest{
		UserID: userID,
		AppID:  appID,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp2)
	}

	resp3, err := GetByApp(context.Background(), &mytype.GetAppLoginRecordsRequest{
		AppID: appID,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp3)
	}

	resp4, err := GetAll(context.Background(), &mytype.GetLoginRecordsRequest{})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp4)
	}
}
