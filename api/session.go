package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/NpoolPlatform/login-door/pkg/response"
	"github.com/NpoolPlatform/login-door/pkg/session"
)

func RefreshSession(w http.ResponseWriter, r *http.Request) {
	// swagger:route Post /v1/refresh/session RefreshSession
	// refresh user's session which maintain user login status.
	// this api can be used when user operate apis.
	// Response:
	//  		200: RefreshSessionResponse
	request := mytype.RefreshSessionRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	info := mytype.LoginSession{
		LoginIP:    r.RemoteAddr,
		LoginTime:  time.Now().Local().String(),
		LoginAgent: r.UserAgent(),
		AppID:      request.AppID,
		UserID:     request.UserID,
		Session:    request.Session,
	}
	resp, err := session.RefreshSession(&info, request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	http.SetCookie(w, &resp)
	response.RespondwithJSON(w, http.StatusOK, mytype.RefreshSessionResponse{
		Info: "refresh session successfully",
	})
}
