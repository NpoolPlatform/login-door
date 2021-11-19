package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/NpoolPlatform/login-door/pkg/cookie"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	myredis "github.com/NpoolPlatform/login-door/pkg/redis"
)

func GenerateSession(size int) (string, error) {
	b := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func RefreshSession(w http.ResponseWriter, request mytype.RefreshSessionRequest) error {
	loginResp, err := myredis.QueryKeyInfo(mytype.LoginKeyword, request.LoginSession)
	if err != nil {
		return err
	}

	appLoginResp, err := myredis.QueryKeyInfo(mytype.LoginKeyword, request.AppLoginSession)
	if err != nil {
		return err
	}

	err = myredis.InsertKeyInfo(mytype.LoginKeyword, request.LoginSession, loginResp, mytype.SessionExpires)
	if err != nil {
		return err
	}

	err = myredis.InsertKeyInfo(mytype.LoginKeyword, request.AppLoginSession, appLoginResp, mytype.SessionExpires)
	if err != nil {
		return err
	}
	err = cookie.SetAllCookie(request.LoginSession, request.AppLoginSession, request.UserID, request.AppID, w)
	return err
}
