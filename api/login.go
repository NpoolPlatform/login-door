package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/NpoolPlatform/login-door/pkg/cookie"
	loginrecord "github.com/NpoolPlatform/login-door/pkg/crud/login-record"
	"github.com/NpoolPlatform/login-door/pkg/location"
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

	loginSession, err := session.GenerateSession(16)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	appLoginSession, err := session.GenerateSession(16)

	infoLogin := mytype.LoginSession{
		LoginIP:    r.RemoteAddr,
		LoginTime:  time.Now().Local().String(),
		LoginAgent: r.UserAgent(),
		Session:    loginSession,
		UserID:     resp.UserBasicInfo.UserID,
	}

	err = myredis.InsertKeyInfo(mytype.LoginKeyword, loginSession, infoLogin, mytype.SessionExpires)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	infoAppLogin := mytype.LoginSession{
		LoginIP:    r.RemoteAddr,
		LoginTime:  time.Now().Local().String(),
		LoginAgent: r.UserAgent(),
		Session:    appLoginSession,
		AppID:      request.AppID,
		UserID:     resp.UserBasicInfo.UserID,
	}
	err = myredis.InsertKeyInfo(mytype.LoginKeyword, appLoginSession, infoAppLogin, mytype.SessionExpires)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respLocation, err := location.GetLocationByIP(r.RemoteAddr)
	if err != nil {
		_, err := loginrecord.Create(context.Background(), &mytype.LoginRecord{
			UserID:    resp.UserBasicInfo.UserID,
			AppID:     request.AppID,
			IP:        r.RemoteAddr,
			LoginTime: uint32(time.Now().Unix()),
		})
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		_, err := loginrecord.Create(context.Background(), &mytype.LoginRecord{
			UserID:    resp.UserBasicInfo.UserID,
			AppID:     request.AppID,
			IP:        r.RemoteAddr,
			Lat:       float64(respLocation.Lat),
			Lon:       float64(respLocation.Lon),
			LoginTime: uint32(time.Now().Unix()),
			Location:  respLocation.Country + "," + respLocation.City,
		})
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	err = cookie.SetAllCookie(r, loginSession, appLoginSession, resp.UserBasicInfo.UserID, request.AppID, w)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp.UserBasicInfo.Password = ""
	response.RespondwithJSON(w, http.StatusOK, mytype.LoginResponse{
		Info:        resp,
		RedirectURL: request.RedirectURL,
	})
}

// GetSSOLogin swagger:route POST /v1/get/sso/login getSSOLogin
// this api get user sso login info.
// Responses:
//			default: getSSOLoginResponse
func GetSSOLogin(w http.ResponseWriter, r *http.Request) {
	request := mytype.GetSSOLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := login.GetSSOLogin(request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	appLoginSession, err := session.GenerateSession(16)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = cookie.SetAllCookie(r, request.Session, appLoginSession, request.UserID, request.AppID, w)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.RespondwithJSON(w, http.StatusOK, resp)
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
		Name:    mytype.AppLoginSessionKey,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
		Path:    "/",
	})
	response.RespondwithJSON(w, http.StatusOK, mytype.LogoutResponse{
		Info: "logout successfully",
	})
}
