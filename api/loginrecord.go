// nolint

package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	loginrecord "github.com/NpoolPlatform/login-door/pkg/crud/login-record"
	"github.com/NpoolPlatform/login-door/pkg/location"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/NpoolPlatform/login-door/pkg/response"
)

// CreateUserLoginRecord swagger:route POST /v1/create/user/login/record createUserLoginRecord
// Create user login history record.
// Responses:
// 			default: createUserLoginRecordResponse
func CreateUserLoginRecord(w http.ResponseWriter, r *http.Request) {
	request := mytype.CreateUserLoginRecordRequest{}
	respLocation, err := location.GetLocationByIP(r.RemoteAddr)
	if err != nil {
		_, err := loginrecord.Create(context.Background(), &mytype.LoginRecord{
			UserID:    request.UserID,
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
			UserID:    request.UserID,
			AppID:     request.AppID,
			IP:        r.RemoteAddr,
			Lat:       float64(respLocation.Lat),
			Lon:       float64(respLocation.Lon),
			LoginTime: uint32(time.Now().Unix()),
			Location:  respLocation.City + ", " + respLocation.Country,
		})
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}
	response.RespondwithJSON(w, http.StatusOK, mytype.CreateUserLoginRecordResponse{
		Info: "create user login history successfully!",
	})
}

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
