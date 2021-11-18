package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/login-door/pkg/cookie"
	"github.com/NpoolPlatform/login-door/pkg/login"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	myredis "github.com/NpoolPlatform/login-door/pkg/redis"
	"github.com/NpoolPlatform/login-door/pkg/response"
	"github.com/NpoolPlatform/login-door/pkg/session"
)

// Login swagger:route POST /v1/login login
// login into system.
// this api implement login by username and password, email or phone verification code login and third party login.
// Responses:
// 			default: loginResponse
func Login(w http.ResponseWriter, r *http.Request) {
	request := mytype.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := login.Login(r, &request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	session, err := session.GenerateSession(16)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	info := mytype.LoginSession{
		LoginIP:    r.RemoteAddr,
		LoginTime:  time.Now().Local().String(),
		LoginAgent: r.UserAgent(),
		Session:    session,
		AppID:      request.AppID,
		UserID:     resp,
	}

	err = myredis.InsertKeyInfo(mytype.LoginKeyword, session, info, mytype.SessionExpires)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	serviceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	cookieDomain := config.GetStringValueWithNameSpace(serviceName, cookie.CookieDomain)

	myCookie := cookie.CreateLoginSessionCookie(session)
	http.SetCookie(w, &myCookie)
	http.SetCookie(w, &http.Cookie{
		Name:    "npool_user_id",
		Value:   resp,
		Path:    "/",
		Domain:  cookieDomain,
		Expires: time.Now().AddDate(0, 0, 1),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "npool_app_id",
		Value:   request.AppID,
		Path:    "/",
		Domain:  cookieDomain,
		Expires: time.Now().AddDate(0, 0, 1),
	})
	http.Redirect(w, r, "/", http.StatusFound)
	response.RespondwithJSON(w, http.StatusOK, mytype.LoginResponse{
		Info: "login successfully",
	})
}

// GetUserLogin swagger:route POST /v1/get/user/login getUserLogin
// this api get user login status and the user's login info stored in redis key.
// Responses:
//      default: getUserLoginResponse
func GetUserLogin(w http.ResponseWriter, r *http.Request) {
	request := mytype.GetUserLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := login.GetUserLogin(request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	response.RespondwithJSON(w, http.StatusOK, resp)
	return
}

// Logout swagger:route POST /v1/logout logout
// this api implement logout from system
// Responses:
//			default: logoutResponse
func Logout(w http.ResponseWriter, r *http.Request) {
	request := mytype.LogoutRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = login.Logout(request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    cookie.LoginSessionCookieName,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
		Path:    "/",
		Domain:  cookie.CookieDomain,
	})
	http.Redirect(w, r, "/", http.StatusFound)
	response.RespondwithJSON(w, http.StatusOK, mytype.LogoutResponse{
		Info: "logout successfully",
	})
}
