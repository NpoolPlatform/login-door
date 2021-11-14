package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NpoolPlatform/login-door/pkg/crud/provider"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/NpoolPlatform/login-door/pkg/response"
)

func AddProvider(w http.ResponseWriter, r *http.Request) {
	request := mytype.AddProviderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("json err is", err)
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	resp, err := provider.Create(context.Background(), &request)
	if err != nil {
		fmt.Println("resp err is", err)
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	response.RespondwithJSON(w, http.StatusOK, &resp)
}

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

func GetAllProviders(w http.ResponseWriter, r *http.Request) {
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
