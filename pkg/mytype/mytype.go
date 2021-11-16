package mytype

import (
	"time"

	"github.com/google/uuid"
)

const (
	LoginKeyword   = "login"
	SessionExpires = 12 * time.Hour
)

type VersionResponse struct {
	Info string `json:"info"`
}

// swagger:parameters add providers
type AddProviderRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	ProviderName string `json:"provider_name"`
	ProviderLogo string `json:"provider_logo,omitempty"`
	ProviderURL  string `json:"provider_url,omitempty"`
}

type ProviderInfo struct {
	ProviderID   uuid.UUID `json:"provider_id"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	ProviderName string    `json:"provider_name"`
	ProviderLogo string    `json:"provider_logo,omitempty"`
	ProviderURL  string    `json:"provider_url,omitempty"`
}

// swagger:response AddProviderResponse
type AddProviderResponse struct {
	Info ProviderInfo `json:"info"`
}

// swagger:parameters UpdateProviderRequest
type UpdateProviderRequest struct {
	Info ProviderInfo `json:"info"`
}

// swagger:response UpdateProviderResponse
type UpdateProviderResponse struct {
	Info ProviderInfo `json:"info"`
}

type PageInfo struct {
	PageIndex int32 `json:"page_index"`
	PageSize  int32 `json:"page_size"`
}

// swagger:parameters GetProvidersRequest
type GetProvidersRequest struct {
	Info PageInfo `json:"info,omitempty"`
}

// swagger:response GetProvidersResponse
type GetProvidersResponse struct {
	Infos []ProviderInfo `json:"infos"`
}

// swagger:parameters DeleteProviderRequest
type DeleteProviderRequest struct {
	ProviderID uuid.UUID `json:"provider_id"`
}

// swagger:response DeleteProviderResponse
type DeleteProviderResponse struct {
	Info string `json:"info"`
}

type LoginSession struct {
	LoginIP    string `json:"login_ip"`
	LoginTime  string `json:"login_time"`
	LoginAgent string `json:"login_agent"`
	Session    string `json:"session"`
	UserID     string `json:"user_id"`
	AppID      string `json:"app_id"`
}

// swagger:parameters LoginRequest
type LoginRequest struct {
	AppID       string `json:"app_id"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	VerifyCode  string `json:"verify_code,omitempty"`
	Provider    string `json:"provider,omitempty"` // provider is provider id
	Code        string `json:"code,omitempty"`
	State       string `json:"state,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`
	Method      string `json:"method,omitempty"`
}

// swagger:response LoginResponse
type LoginResponse struct {
	Info string `json:"info"`
}

// swagger:parameters GetUserLoginRequest
type GetUserLoginRequest struct {
	Session string `json:"session"`
	UserID  string `json:"user_id"`
}

// swagger:response GetUserLoginResponse
type GetUserLoginResponse struct {
	Info LoginSession `json:"info"`
}

// swagger:parameters RefreshSessionRequest
type RefreshSessionRequest struct {
	Session string `json:"session"`
	UserID  string `json:"user_id"`
	AppID   string `json:"app_id"`
}

// swagger:response RefreshSessionResponse
type RefreshSessionResponse struct {
	Info string `json:"info"`
}

// swagger:parameters LogoutRequest
type LogoutRequest struct {
	Session string `json:"session"`
	UserID  string `json:"user_id"`
}

// swagger:response LogoutResponse
type LogoutResponse struct {
	Info string `json:"info"`
}
