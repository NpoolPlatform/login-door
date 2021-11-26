package login

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/login-door/pkg/grpc"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	testinit "github.com/NpoolPlatform/login-door/pkg/test-init"
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

func TestLogin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appInfo, err := grpc.CreaeteApp()
	if assert.Nil(t, err) {
		assert.NotNil(t, appInfo)
	}
	fmt.Println("app info is", appInfo)

	userInfo, err := grpc.CreateTestUser(appInfo.Info.ID)
	if assert.Nil(t, err) {
		assert.NotNil(t, userInfo)
	}
	fmt.Println("user info is", userInfo)

	resp, err := ByUsername(&mytype.LoginRequest{
		AppID:    appInfo.Info.ID,
		Username: userInfo.Info.Username,
		Password: "12345679",
	})
	fmt.Println(err)
	if assert.Nil(t, err) {
		assert.Equal(t, userInfo.Info.UserID, resp.UserBasicInfo.UserID)
	}
}
