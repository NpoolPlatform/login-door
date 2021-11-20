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

func RefreshSession(r *http.Request, w http.ResponseWriter, request mytype.RefreshSessionRequest) error {
	loginResp, err := myredis.QueryKeyInfo(mytype.LoginKeyword, request.Session)
	if err != nil {
		return err
	}

	appLoginResp, err := myredis.QueryKeyInfo(mytype.LoginKeyword, request.AppSession)
	if err != nil {
		return err
	}

	err = myredis.InsertKeyInfo(mytype.LoginKeyword, request.Session, loginResp, mytype.SessionExpires)
	if err != nil {
		return err
	}

	err = myredis.InsertKeyInfo(mytype.LoginKeyword, request.AppSession, appLoginResp, mytype.SessionExpires)
	if err != nil {
		return err
	}
	err = cookie.SetAllCookie(r, request.Session, request.AppSession, request.UserID, request.AppID, w)
	return err
}
