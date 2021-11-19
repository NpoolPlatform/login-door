package cookie

import (
	"net/http"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
)

const CookieDomain = "cookie_domain"

func getCookieDomain() string {
	serviceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	cookieDomain := config.GetStringValueWithNameSpace(serviceName, CookieDomain)
	return cookieDomain
}

func SetAllCookie(loginSession, appLoginSession, userID, appID string, w http.ResponseWriter) error {
	cookieDomain := getCookieDomain()
	http.SetCookie(w, &http.Cookie{
		Name:    mytype.AppLoginSessionKey,
		Value:   appLoginSession,
		Path:    "/",
		Expires: time.Now().AddDate(0, 0, 1),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    mytype.LoginSessionKey,
		Value:   loginSession,
		Path:    "/",
		Domain:  cookieDomain,
		Expires: time.Now().AddDate(0, 0, 1),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    mytype.UserIDKey,
		Value:   userID,
		Path:    "/",
		Domain:  cookieDomain,
		Expires: time.Now().AddDate(0, 0, 1),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    mytype.AppIDKey,
		Value:   appID,
		Path:    "/",
		Domain:  cookieDomain,
		Expires: time.Now().AddDate(0, 0, 1),
	})

	return nil
}
