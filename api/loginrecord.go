// nolint

package api

import (
	"context"
	"encoding/json"
	"net/http"

	loginrecord "github.com/NpoolPlatform/login-door/pkg/crud/login-record"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/NpoolPlatform/login-door/pkg/response"
)

// GetUserLoginRecords swagger:route POST /v1/get/user/login/records getUserLoginRecords
// get user login history records.
// Responses:
// 			default: getUserLoginRecordsResponse
func GetUserLoginRecords(w http.ResponseWriter, r *http.Request) {
	request := mytype.GetUserLoginRecordsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := loginrecord.GetByUser(context.Background(), &request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.RespondwithJSON(w, http.StatusOK, resp)
}

// GetAppLoginRecords swagger:route POST /v1/get/app/login/records getAppLoginRecords
// get app users login history records.
// Responses:
// 			default: getAppLoginRecordsResponse
func GetAppLoginRecords(w http.ResponseWriter, r *http.Request) {
	request := mytype.GetAppLoginRecordsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := loginrecord.GetByApp(context.Background(), &request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.RespondwithJSON(w, http.StatusOK, resp)
}

// GetAllLoginRecords swagger:route POST /v1/get/all/login/records getAllLoginRecords
// get all login history records.
// Responses:
// 			default: getAllLoginRecordsResponse
func GetAllLoginRecords(w http.ResponseWriter, r *http.Request) {
	request := mytype.GetLoginRecordsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := loginrecord.GetAll(context.Background(), &request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.RespondwithJSON(w, http.StatusOK, resp)
}
