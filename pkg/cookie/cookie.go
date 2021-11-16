package cookie

import (
	"net/http"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
)

const (
	LoginSessionCookieName = "npool_login_session_id"
	CookieDomain           = "cookie_domain"
)

func CreateLoginSessionCookie(sessionID string) http.Cookie {
	serviceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	cookieDomain := config.GetStringValueWithNameSpace(serviceName, CookieDomain)
	cookie := http.Cookie{
		Name:    LoginSessionCookieName,
		Value:   sessionID,
		Domain:  cookieDomain,
		Expires: time.Now().AddDate(0, 0, 1), // set cookie expires after one day
		Path:    "/",
	}
	return cookie
}
