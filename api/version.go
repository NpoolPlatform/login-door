// +build !codeanalysis

package api

import (
	"net/http"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/login-door/pkg/response"
	"github.com/NpoolPlatform/login-door/pkg/version"
)

func Version(w http.ResponseWriter, r *http.Request) {
	resp, err := version.Version()
	if err != nil {
		logger.Sugar().Errorw("[Version] get service version error: %w", err)
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	response.RespondwithJSON(w, http.StatusOK, resp)
}
