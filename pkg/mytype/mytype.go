package mytype

import (
	"time"

	pbapp "github.com/NpoolPlatform/application-management/message/npool"
	pbuser "github.com/NpoolPlatform/user-management/message/npool"
	"github.com/google/uuid"
)

const (
	LoginKeyword       = "login"
	SessionExpires     = 24 * time.Hour
	AppIDKey           = "AppID"
	UserIDKey          = "UserID"
	LoginSessionKey    = "Session"
	AppLoginSessionKey = "AppSession"
)

type VersionResponse struct {
	Info string
}

type LoginRecord struct {
	ID        string
	UserID    string
	AppID     string
	LoginTime uint32
	IP        string
	Location  string
	Lat       float64
	Lon       float64
	Timezone  string
}

// swagger:parameters addProvider
type AddProviderRequest struct {
	// ClientID
	//
	// in: body
	// required: true
	ClientID string
	// ClientSecret
	//
	// in: body
	// required: true
	ClientSecret string
	// ProviderName, such as: github, google, wechat and so on
	//
	// in: body
	// required: true
	ProviderName string
	// ProviderLogo
	ProviderLogo string
	// ProviderURL
	ProviderURL string
	AppID       string
	UserID      string
}

// swagger:model providerInfo
type ProviderInfo struct {
	ProviderID   uuid.UUID
	ClientID     string
	ClientSecret string
	ProviderName string
	ProviderLogo string
	ProviderURL  string
}

// swagger:response addProviderResponse
type AddProviderResponse struct {
	Info ProviderInfo
}

// swagger:parameters updateProvider
type UpdateProviderRequest struct {
	Info   ProviderInfo
	UserID string
	AppID  string
}

// swagger:response updateProviderResponse
type UpdateProviderResponse struct {
	Info ProviderInfo
}

type PageInfo struct {
	PageIndex int32
	PageSize  int32
}

// swagger:parameters getProviders
type GetAllProvidersRequest struct {
	// page info
	Info   PageInfo
	AppID  string
	UserID string
}

// swagger:response getProvidersResponse
type GetAllProvidersResponse struct {
	Infos []ProviderInfo
}

// swagger:parameters deleteProvider
type DeleteProviderRequest struct {
	// required: true
	ProviderID uuid.UUID
	AppID      string
	UserID     string
}

// swagger:response deleteProviderResponse
type DeleteProviderResponse struct {
	Info string
}

// swagger:model loginSession
type LoginSession struct {
	// login from ip
	LoginIP string
	// when user login
	LoginTime string
	// agent user login from
	LoginAgent string
	// user login session
	Session string
	// user id
	UserID string
	// app id, which app user login
	AppID string
}

// swagger:parameters login
type LoginRequest struct {
	// app id, which app user login
	// in: body
	// required: true
	AppID    string
	Username string
	Password string
	Email    string
	Phone    string
	// email or phone verify code
	VerifyCode string
	// google recaptcha response
	GoogleRecaptchaResponse string
	// Provider id
	Provider string
	// code is returned by provider after user authenticate from provider
	Code string
	// state is returned by provider after user authenticate from provider
	State string
	// redirect url tell provider which callback it need to return after get user info
	RedirectURL string
	Method      string
}

type UserDetail struct {
	UserBasicInfo *pbuser.UserBasicInfo
	UserAppInfo   *pbapp.ApplicationUserDetail
}

// swagger:response loginResponse
type LoginResponse struct {
	Info        *UserDetail
	RedirectURL string
}

// swagger:parameters getUserLogin
type GetUserLoginRequest struct {
	// session stored in cookie
	AppSession string
	// user id stored in cookie
	UserID string
	AppID  string
}

// swagger:response getUserLoginResponse
type GetUserLoginResponse struct {
	Info LoginSession
}

// swagger:parameters getSSOLogin
type GetSSOLoginRequest struct {
	Session string
	UserID  string
	AppID   string
}

// swagger:response getSSOLoginResponse
type GetSSOLoginResponse struct {
	Info LoginSession
}

// swagger:parameters refreshSession
type RefreshSessionRequest struct {
	AppSession string
	Session    string
	UserID     string
	AppID      string
}

// swagger:response refreshSessionResponse
type RefreshSessionResponse struct {
	Info string
}

// swagger:parameters logout
type LogoutRequest struct {
	// session stored in cookie
	AppSession string
	// user id stored in cookie
	UserID string
	AppID  string
}

// swagger:response logoutResponse
type LogoutResponse struct {
	Info string
}

// swagger:parameters getUserLoginRecords
type GetUserLoginRecordsRequest struct {
	UserID string
	AppID  string
}

// swagger:response getUserLoginRecordsResponse
type GetUserLoginRecordsResponse struct {
	Infos []*LoginRecord
}

// swagger:parameters getAppLoginRecords
type GetAppLoginRecordsRequest struct {
	AppID  string
	UserID string
}

// swagger:response getAppLoginRecordsResponse
type GetAppLoginRecordsResponse struct {
	Infos []*LoginRecord
}

// swagger:parameters getAllLoginRecords
type GetLoginRecordsRequest struct {
	AppID  string
	UserID string
}

// swagger:response getAllLoginRecordResposne
type GetLoginRecordsResponse struct {
	Infos []*LoginRecord
}
