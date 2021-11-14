package provider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"

	"github.com/NpoolPlatform/login-door/pkg/mytype"
	testinit "github.com/NpoolPlatform/login-door/pkg/test-init"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestProviderCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	providerInfo := mytype.ProviderInfo{
		ClientID:     randstr.Hex(10),
		ClientSecret: randstr.Hex(20),
		ProviderName: "test",
	}

	resp, err := Create(context.Background(), &mytype.AddProviderRequest{
		ClientID:     providerInfo.ClientID,
		ClientSecret: providerInfo.ClientSecret,
		ProviderName: providerInfo.ProviderName,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ProviderID, uuid.UUID{})
		assert.Equal(t, resp.Info.ClientID, providerInfo.ClientID)
		assert.Equal(t, resp.Info.ClientSecret, providerInfo.ClientSecret)
		assert.Equal(t, resp.Info.ProviderName, providerInfo.ProviderName)
		providerInfo.ProviderID = resp.Info.ProviderID
	}

	providerInfo.ProviderLogo = "test"
	resp1, err := Update(context.Background(), &mytype.UpdateProviderRequest{
		Info: providerInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ProviderID, providerInfo.ProviderID)
		assert.Equal(t, resp1.Info.ClientID, providerInfo.ClientID)
		assert.Equal(t, resp1.Info.ClientSecret, providerInfo.ClientSecret)
		assert.Equal(t, resp1.Info.ProviderName, providerInfo.ProviderName)
		assert.Equal(t, resp1.Info.ProviderLogo, providerInfo.ProviderLogo)
	}

	resp2, err := GetAll(context.Background(), &mytype.GetProvidersRequest{})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Infos[0].ProviderID, providerInfo.ProviderID)
		assert.Equal(t, resp2.Infos[0].ClientID, providerInfo.ClientID)
		assert.Equal(t, resp2.Infos[0].ClientSecret, providerInfo.ClientSecret)
		assert.Equal(t, resp2.Infos[0].ProviderName, providerInfo.ProviderName)
		assert.Equal(t, resp2.Infos[0].ProviderLogo, providerInfo.ProviderLogo)
	}

	_, err = Delete(context.Background(), &mytype.DeleteProviderRequest{
		ProviderID: providerInfo.ProviderID,
	})
	assert.Nil(t, err)
}
