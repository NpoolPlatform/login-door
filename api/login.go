package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NpoolPlatform/login-door/pkg/cookie"
	"github.com/NpoolPlatform/login-door/pkg/login"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	myredis "github.com/NpoolPlatform/login-door/pkg/redis"
	"github.com/NpoolPlatform/login-door/pkg/response"
	"github.com/NpoolPlatform/login-door/pkg/session"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// swagger:route Post /v1/login Login
	// login into system.
	// this api implement login by username and password, email or phone verification code login and third party login.
	// Response:
	// 			200: LoginResponse
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

	myCookie := cookie.CreateLoginSessionCookie(session)
	http.SetCookie(w, &myCookie)
	http.Redirect(w, r, "/", http.StatusFound)
	response.RespondwithJSON(w, http.StatusOK, mytype.LoginResponse{
		Info: "login successfully",
	})
}

func GetUserLogin(w http.ResponseWriter, r *http.Request) {
	// swagger:route Post /v1/get/user/login
	// this api get user login status and the user's login info stored in redis key.
	// Response:
	//      200: GetUserLoginResponse
	request := mytype.GetUserLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := login.GetUserLogin(request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	response.RespondwithJSON(w, http.StatusOK, resp)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// swagger:route Post /v1/logout
	// this api implement logout from system
	// Response:
	//			200: LogoutResponse
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
