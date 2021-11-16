package provider

import (
	"context"
	"time"

	"github.com/NpoolPlatform/login-door/pkg/db"
	"github.com/NpoolPlatform/login-door/pkg/db/ent"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/provider"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToProviderInfo(row *ent.Provider) mytype.ProviderInfo {
	return mytype.ProviderInfo{
		ProviderID:   row.ID,
		ClientID:     row.ClientID,
		ClientSecret: row.ClientSecret,
		ProviderName: row.ProviderName,
		ProviderLogo: row.ProviderLogo,
		ProviderURL:  row.ProviderURL,
	}
}

func Create(ctx context.Context, in *mytype.AddProviderRequest) (*mytype.AddProviderResponse, error) {
	info, err := db.Client().
		Provider.
		Create().
		SetClientID(in.ClientID).
		SetClientSecret(in.ClientSecret).
		SetProviderName(in.ProviderName).
		SetProviderLogo(in.ProviderLogo).
		SetProviderURL(in.ProviderURL).
		Save(ctx)
	if err != nil {
		return &mytype.AddProviderResponse{}, xerrors.Errorf("fail to create provider info: %v", err)
	}

	return &mytype.AddProviderResponse{
		Info: dbRowToProviderInfo(info),
	}, nil
}

func Update(ctx context.Context, in *mytype.UpdateProviderRequest) (*mytype.UpdateProviderResponse, error) {
	_, err := db.Client().
		Provider.
		Query().
		Where(
			provider.And(
				provider.ID(in.Info.ProviderID),
				provider.DeleteAt(0),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query provider info: %v", err)
	}

	info, err := db.Client().
		Provider.
		UpdateOneID(in.Info.ProviderID).
		SetClientID(in.Info.ClientID).
		SetClientSecret(in.Info.ClientSecret).
		SetProviderName(in.Info.ProviderName).
		SetProviderLogo(in.Info.ProviderLogo).
		SetProviderURL(in.Info.ProviderURL).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update provider: %v", err)
	}

	return &mytype.UpdateProviderResponse{
		Info: dbRowToProviderInfo(info),
	}, nil
}

func Get(ctx context.Context, providerID string) (mytype.ProviderInfo, error) {
	id, err := uuid.Parse(providerID)
	if err != nil {
		return mytype.ProviderInfo{}, xerrors.Errorf("invalid provider id: %v", err)
	}

	info, err := db.Client().
		Provider.
		Query().
		Where(
			provider.And(
				provider.DeleteAt(0),
				provider.ID(id),
			),
		).Only(ctx)
	if err != nil {
		return mytype.ProviderInfo{}, xerrors.Errorf("fail to get provider info: %v", err)
	}
	return dbRowToProviderInfo(info), nil
}

func GetAll(ctx context.Context, in *mytype.GetProvidersRequest) (*mytype.GetProvidersResponse, error) {
	infos, err := db.Client().
		Provider.
		Query().
		Where(
			provider.And(
				provider.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get all providers: %v", err)
	}

	response := []mytype.ProviderInfo{}
	for _, info := range infos {
		response = append(response, dbRowToProviderInfo(info))
	}
	return &mytype.GetProvidersResponse{
		Infos: response,
	}, nil
}

func Delete(ctx context.Context, in *mytype.DeleteProviderRequest) (*mytype.DeleteProviderResponse, error) {
	_, err := db.Client().
		Provider.
		UpdateOneID(in.ProviderID).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete provider: %v", err)
	}

	return &mytype.DeleteProviderResponse{
		Info: "delete provider successfully",
	}, nil
}
