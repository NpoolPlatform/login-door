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

func RefreshSession(info *mytype.LoginSession, request mytype.RefreshSessionRequest) (http.Cookie, error) {
	err := myredis.InsertKeyInfo(mytype.LoginKeyword, request.Session, info, mytype.SessionExpires)
	if err != nil {
		return http.Cookie{}, err
	}
	myCookie := cookie.CreateLoginSessionCookie(request.Session)
	return myCookie, nil
}
