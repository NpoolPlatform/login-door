package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/NpoolPlatform/login-door/pkg/crud/provider"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/NpoolPlatform/login-door/pkg/response"
)

// AddProvider swagger:route POST /v1/add/provider addProvider
// add provider(third party) login info.
// Responses:
//      default: addProviderResponse
func AddProvider(w http.ResponseWriter, r *http.Request) {
	request := mytype.AddProviderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	resp, err := provider.Create(context.Background(), &request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	response.RespondwithJSON(w, http.StatusOK, &resp)
}

// UpdateProvider swagger:route POST /v1/update/provider updateProvider
// update provider(third party) login info.
// Responses:
//      default: updateProviderResponse
func UpdateProvider(w http.ResponseWriter, r *http.Request) {
	request := mytype.UpdateProviderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	resp, err := provider.Update(context.Background(), &request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	response.RespondwithJSON(w, http.StatusOK, &resp)
}

// GetAllProviders swagger:route POST /v1/get/all/providers getProviders
// get all providers(third party) login infos.
// Responses:
//      default: getProvidersResponse
func GetAllProviders(w http.ResponseWriter, r *http.Request) {
	request := mytype.GetAllProvidersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	resp, err := provider.GetAll(context.Background(), &request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	response.RespondwithJSON(w, http.StatusOK, &resp)
}

// DeletProvider swagger:route POST /v1/delete/provider deleteProvider
// delete provider(third party) login info.
// Responses:
//      default: deleteProviderResponse
func DeleteProvider(w http.ResponseWriter, r *http.Request) {
	request := mytype.DeleteProviderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	resp, err := provider.Delete(context.Background(), &request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	response.RespondwithJSON(w, http.StatusOK, &resp)
}
