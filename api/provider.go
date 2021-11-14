package api

import (
	"encoding/json"
	"net/http"

	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/NpoolPlatform/login-door/pkg/response"
)

func AddProvider(w http.ResponseWriter, r *http.Request) {
	request := mytype.AddProviderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

}
