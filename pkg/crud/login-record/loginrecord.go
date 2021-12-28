package loginrecord

import (
	"context"
	"time"

	"github.com/NpoolPlatform/login-door/pkg/db"
	"github.com/NpoolPlatform/login-door/pkg/db/ent"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/loginrecord"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToInfo(row *ent.LoginRecord) *mytype.LoginRecord {
	return &mytype.LoginRecord{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		UserID:    row.UserID.String(),
		IP:        row.IP,
		Lat:       row.Lat,
		Lon:       row.Lon,
		LoginTime: row.LoginTime,
		Location:  row.Location,
		Timezone:  row.Timezone,
	}
}

func Create(ctx context.Context, in *mytype.LoginRecord) (*mytype.LoginRecord, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	appID, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		LoginRecord.
		Create().
		SetAppID(appID).
		SetUserID(userID).
		SetIP(in.IP).
		SetLat(in.Lat).
		SetLon(in.Lon).
		SetLoginTime(in.LoginTime).
		SetLocation(in.Location).
		SetTimezone(in.Timezone).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create login record")
	}

	return dbRowToInfo(info), nil
}

func GetByUser(ctx context.Context, in *mytype.GetUserLoginRecordsRequest) (*mytype.GetUserLoginRecordsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	app, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		LoginRecord.
		Query().
		Where(
			loginrecord.And(
				loginrecord.AppID(app),
				loginrecord.UserID(user),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get user login record: %v", err)
	}

	response := []*mytype.LoginRecord{}
	for _, info := range infos {
		response = append(response, dbRowToInfo(info))
	}

	return &mytype.GetUserLoginRecordsResponse{
		Infos: response,
	}, nil
}

func GetByApp(ctx context.Context, in *mytype.GetAppLoginRecordsRequest) (*mytype.GetAppLoginRecordsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	appID, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		LoginRecord.
		Query().
		Where(
			loginrecord.And(
				loginrecord.AppID(appID),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get login records by app: %v", err)
	}

	response := []*mytype.LoginRecord{}
	for _, info := range infos {
		response = append(response, dbRowToInfo(info))
	}

	return &mytype.GetAppLoginRecordsResponse{
		Infos: response,
	}, nil
}

func GetAll(ctx context.Context, in *mytype.GetLoginRecordsRequest) (*mytype.GetLoginRecordsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		LoginRecord.
		Query().All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get all login records: %v", err)
	}

	response := []*mytype.LoginRecord{}
	for _, info := range infos {
		response = append(response, dbRowToInfo(info))
	}

	return &mytype.GetLoginRecordsResponse{
		Infos: response,
	}, nil
}
