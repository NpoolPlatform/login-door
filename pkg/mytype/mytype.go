package mytype

import "github.com/google/uuid"

type VersionResponse struct {
	Info string `json:"info"`
}

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

type AddProviderResponse struct {
	Info ProviderInfo `json:"info"`
}

type UpdateProviderRequest struct {
	Info ProviderInfo `json:"info"`
}

type UpdateProviderResponse struct {
	Info ProviderInfo `json:"info"`
}

type PageInfo struct {
	PageIndex int32 `json:"page_index"`
	PageSize  int32 `json:"page_size"`
}

type GetProvidersRequest struct {
	Info PageInfo `json:"info,omitempty"`
}

type GetProvidersResponse struct {
	Infos []ProviderInfo `json:"infos"`
}

type DeleteProviderRequest struct {
	ProviderID uuid.UUID `json:"provider_id"`
}

type DeleteProviderResponse struct {
	Info string `json:"info"`
}
