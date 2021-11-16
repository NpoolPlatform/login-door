package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/NpoolPlatform/login-door/pkg/crud/provider"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/NpoolPlatform/login-door/pkg/response"
)

func AddProvider(w http.ResponseWriter, r *http.Request) {
	// swagger:route Post /v1/add/provider AddProvider
	// add provider(third party) login info.
	// Response:
	//      200: AddProviderResponse
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

func UpdateProvider(w http.ResponseWriter, r *http.Request) {
	// swagger:route Post /v1/update/provider UpdateProvider
	// update provider(third party) login info.
	// Response:
	//      200: UpdateProviderResponse
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

func GetAllProviders(w http.ResponseWriter, r *http.Request) {
	// swagger:route Post /v1/get/all/providers GetAllProviders
	// get all providers(third party) login infos.
	// Response:
	//      200: GetAllProvidersResponse
	request := mytype.GetProvidersRequest{}
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

func DeleteProvider(w http.ResponseWriter, r *http.Request) {
	// swagger:route Post /v1/delete/provider DeleteProvider
	// delete provider(third party) login info.
	// Response:
	//      200: DeleteProviderResponse
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
