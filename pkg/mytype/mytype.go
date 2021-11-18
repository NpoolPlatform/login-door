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

// swagger:parameters addProvider
type AddProviderRequest struct {
	// ClientID
	//
	// in: body
	// required: true
	ClientID string `json:"client_id"`
	// ClientSecret
	//
	// in: body
	// required: true
	ClientSecret string `json:"client_secret"`
	// ProviderName, such as: github, google, wechat and so on
	//
	// in: body
	// required: true
	ProviderName string `json:"provider_name"`
	// ProviderLogo
	ProviderLogo string `json:"provider_logo,omitempty"`
	// ProviderURL
	ProviderURL string `json:"provider_url,omitempty"`
}

// swagger:model providerInfo
type ProviderInfo struct {
	ProviderID   uuid.UUID `json:"provider_id"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	ProviderName string    `json:"provider_name"`
	ProviderLogo string    `json:"provider_logo,omitempty"`
	ProviderURL  string    `json:"provider_url,omitempty"`
}

// swagger:response addProviderResponse
type AddProviderResponse struct {
	Info ProviderInfo `json:"info"`
}

// swagger:parameters updateProvider
type UpdateProviderRequest struct {
	Info ProviderInfo `json:"info"`
}

// swagger:response updateProviderResponse
type UpdateProviderResponse struct {
	Info ProviderInfo `json:"info"`
}

type PageInfo struct {
	PageIndex int32 `json:"page_index"`
	PageSize  int32 `json:"page_size"`
}

// swagger:parameters getProviders
type GetAllProvidersRequest struct {
	// page info
	Info PageInfo `json:"info,omitempty"`
}

// swagger:response getProvidersResponse
type GetAllProvidersResponse struct {
	Infos []ProviderInfo `json:"infos"`
}

// swagger:parameters deleteProvider
type DeleteProviderRequest struct {
	// required: true
	ProviderID uuid.UUID `json:"provider_id"`
}

// swagger:response deleteProviderResponse
type DeleteProviderResponse struct {
	Info string `json:"info"`
}

// swagger:model loginSession
type LoginSession struct {
	// login from ip
	LoginIP string `json:"login_ip"`
	// when user login
	LoginTime string `json:"login_time"`
	// agent user login from
	LoginAgent string `json:"login_agent"`
	// user login session
	Session string `json:"session"`
	// user id
	UserID string `json:"user_id"`
	// app id, which app user login
	AppID string `json:"app_id"`
}

// swagger:parameters login
type LoginRequest struct {
	// app id, which app user login
	// in: body
	// required: true
	AppID    string `json:"app_id"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	// email or phone verify code
	VerifyCode string `json:"verify_code,omitempty"`
	// Provider id
	Provider string `json:"provider,omitempty"`
	// code is returned by provider after user authenticate from provider
	Code string `json:"code,omitempty"`
	// state is returned by provider after user authenticate from provider
	State string `json:"state,omitempty"`
	// redirect url tell provider which callback it need to return after get user info
	RedirectURL string `json:"redirect_url,omitempty"`
	Method      string `json:"method,omitempty"`
}

// swagger:response loginResponse
type LoginResponse struct {
	Info string `json:"info"`
}

// swagger:parameters getUserLogin
type GetUserLoginRequest struct {
	// session stored in cookie
	Session string `json:"session"`
	// user id stored in cookie
	UserID string `json:"user_id"`
}

// swagger:response getUserLoginResponse
type GetUserLoginResponse struct {
	Info LoginSession `json:"info"`
}

// swagger:parameters refreshSession
type RefreshSessionRequest struct {
	Session string `json:"session"`
	UserID  string `json:"user_id"`
	AppID   string `json:"app_id"`
}

// swagger:response refreshSessionResponse
type RefreshSessionResponse struct {
	Info string `json:"info"`
}

// swagger:parameters logout
type LogoutRequest struct {
	// session stored in cookie
	Session string `json:"session"`
	// user id stored in cookie
	UserID string `json:"user_id"`
}

// swagger:response logoutResponse
type LogoutResponse struct {
	Info string `json:"info"`
}
