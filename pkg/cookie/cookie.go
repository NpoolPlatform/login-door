package cookie

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/NpoolPlatform/login-door/pkg/mytype"
)

func getCookieDomain(r *http.Request) string {
	host := strings.Split(r.Host, ":")[0]
	if !strings.Contains(host, ".") {
		return ""
	}

	splitHost := strings.Split(host, ".")
	for _, s := range splitHost {
		if _, err := strconv.ParseInt(s, 10, 64); err != nil {
			return ""
		}
	}
	return strings.Join(splitHost[len(splitHost)-2:], ".")
}

func SetAllCookie(r *http.Request, loginSession, appLoginSession, userID string, w http.ResponseWriter) error {
	cookieDomain := getCookieDomain(r)
	if cookieDomain != "" {
		http.SetCookie(w, &http.Cookie{
			Name:     mytype.LoginSessionKey,
			Value:    loginSession,
			Path:     "/",
			SameSite: http.SameSite(2),
			Domain:   cookieDomain,
			Expires:  time.Now().AddDate(0, 0, 1),
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     mytype.LoginSessionKey,
			Value:    loginSession,
			Path:     "/",
			SameSite: http.SameSite(2),
			Expires:  time.Now().AddDate(0, 0, 1),
		})
	}

	http.SetCookie(w, &http.Cookie{
		Name:     mytype.AppLoginSessionKey,
		Value:    appLoginSession,
		Path:     "/",
		SameSite: http.SameSite(2),
		Expires:  time.Now().AddDate(0, 0, 1),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     mytype.UserIDKey,
		Value:    userID,
		Path:     "/",
		SameSite: http.SameSite(2),
		Expires:  time.Now().AddDate(0, 0, 1),
	})

	return nil
}
