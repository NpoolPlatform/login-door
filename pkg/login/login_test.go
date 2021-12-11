package login

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/login-door/pkg/test-init"
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
}
