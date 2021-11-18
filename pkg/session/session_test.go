package session

import (
	"fmt"
	"os"
	"strconv"
	"testing"

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

func TestSession(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	session, err := GenerateSession(16)
	if assert.Nil(t, err) {
		assert.NotNil(t, session)
	}

	userID := uuid.New().String()
	appID := uuid.New().String()
	info := &mytype.LoginSession{
		LoginIP:    "test",
		LoginTime:  "test",
		LoginAgent: "test",
		Session:    session,
		UserID:     userID,
		AppID:      appID,
	}
	cookie, err := RefreshSession(info, mytype.RefreshSessionRequest{
		Session: session,
		AppID:   appID,
		UserID:  userID,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, cookie)
	}
}
