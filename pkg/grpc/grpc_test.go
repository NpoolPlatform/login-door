package grpc

import (
	"fmt"
	"os"
	"strconv"
	"testing"

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

func TestGrpc(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	err := VerifyCode("crazyzplzpl@163.com", "12345")
	assert.NotNil(t, err)

	username := "test" + uuid.New().String()
	_, err = QueryUserExist(username, "")
	assert.NotNil(t, err)

	_, err = QueryUserByUserProviderID("", "")
	assert.NotNil(t, err)
}
