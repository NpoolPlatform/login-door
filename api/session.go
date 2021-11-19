package api

import (
	"encoding/json"
	"net/http"

	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/NpoolPlatform/login-door/pkg/response"
	"github.com/NpoolPlatform/login-door/pkg/session"
)

// RefreshSession swagger:route Post /v1/refresh/session refreshSession
// refresh user's session which maintain user login status.
// this api can be used when user operate apis.
// Responses:
//  		default: refreshSessionResponse
func RefreshSession(w http.ResponseWriter, r *http.Request) {
	request := mytype.RefreshSessionRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = session.RefreshSession(r, w, request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.RespondwithJSON(w, http.StatusOK, mytype.RefreshSessionResponse{
		Info: "refresh session successfully",
	})
}
