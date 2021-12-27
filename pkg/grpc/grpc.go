package grpc

import (
	"context"
	"encoding/json"
	"time"

	pbapplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pbuser "github.com/NpoolPlatform/user-management/message/npool"
	userconst "github.com/NpoolPlatform/user-management/pkg/message/const"
	pbverification "github.com/NpoolPlatform/verification-door/message/npool"
	verificationconst "github.com/NpoolPlatform/verification-door/pkg/message/const"
	"github.com/casbin/casdoor/idp"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	grpcTimeout = 5 * time.Second
)

func VerifyCode(ctx context.Context, param, code string) error {
	conn, err := mygrpc.GetGRPCConn(verificationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbverification.NewVerificationDoorClient(conn)
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	_, err = client.VerifyCode(ctx, &pbverification.VerifyCodeRequest{
		Param: param,
		Code:  code,
	})
	if err != nil {
		return err
	}
	return nil
}

func VerifyGoogleRecaptcha(ctx context.Context, response string) error {
	conn, err := mygrpc.GetGRPCConn(verificationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbverification.NewVerificationDoorClient(conn)
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()
	resp, err := client.VerifyGoogleRecaptcha(ctx, &pbverification.VerifyGoogleRecaptchaRequest{
		Response: response,
	})
	if err != nil {
		return err
	}

	if !resp.Info {
		return xerrors.Errorf("verify google recaptcha wrong")
	}

	return nil
}

func CreateTestUser(ctx context.Context, appID string) (*pbuser.SignupResponse, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()
	resp, err := client.SignUp(ctx, &pbuser.SignupRequest{
		AppID:    appID,
		Username: uuid.New().String(),
		Password: "12345679",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CreateUser(ctx context.Context, appID, providerID string, providerUserInfo *idp.UserInfo) (*pbuser.BindThirdPartyResponse, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()
	resp, err := client.AddUser(ctx, &pbuser.AddUserRequest{
		AppID: appID,
		UserInfo: &pbuser.UserBasicInfo{
			DisplayName:  providerUserInfo.DisplayName,
			Avatar:       providerUserInfo.AvatarUrl,
			EmailAddress: providerUserInfo.Email,
			Username:     uuid.New().String(),
		},
	})
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(providerUserInfo)
	if err != nil {
		return nil, xerrors.Errorf("fail to marshal provider user info: %v", err)
	}
	respBind, err := client.BindThirdParty(ctx, &pbuser.BindThirdPartyRequest{
		UserID:           resp.Info.UserID,
		AppID:            appID,
		ProviderID:       providerID,
		ProviderUserID:   providerUserInfo.Id,
		UserProviderInfo: string(b),
	})
	if err != nil {
		return nil, err
	}

	return respBind, nil
}

func QueryUserExist(ctx context.Context, username, password string) (*pbuser.UserBasicInfo, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()
	resp, err := client.QueryUserExist(ctx, &pbuser.QueryUserExistRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return resp.Info, nil
}

func QueryUserByUserProviderID(ctx context.Context, providerID, userProviderID string) (*pbuser.UserBasicInfo, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()
	resp, err := client.QueryUserByUserProviderID(ctx, &pbuser.QueryUserByUserProviderIDRequest{
		ProviderID:     providerID,
		ProviderUserID: userProviderID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Info.UserBasicInfo, nil
}

func QueryUserFrozen(ctx context.Context, userID string) error {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()
	resp, err := client.QueryUserFrozen(ctx, &pbuser.QueryUserFrozenRequest{
		UserID: userID,
	})
	if err != nil {
		return nil
	}
	if resp.Info != nil {
		return xerrors.Errorf("user has been frozen")
	}
	return nil
}

func CreaeteApp(ctx context.Context) (*pbapplication.CreateApplicationResponse, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbapplication.NewApplicationManagementClient(conn)
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()
	resp, err := client.CreateApplication(ctx, &pbapplication.CreateApplicationRequest{
		Info: &pbapplication.ApplicationInfo{
			ApplicationName:  uuid.New().String(),
			ApplicationOwner: uuid.New().String(),
		},
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetApplication(ctx context.Context, appID string) (*pbapplication.GetApplicationResponse, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbapplication.NewApplicationManagementClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()
	resp, err := client.GetApplication(ctx, &pbapplication.GetApplicationRequest{
		AppID: appID,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func QueryUserInApplication(ctx context.Context, userID, appID string) error {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbapplication.NewApplicationManagementClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()
	_, err = client.GetUserFromApplication(ctx, &pbapplication.GetUserFromApplicationRequest{
		AppID:  appID,
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserDetail(ctx context.Context, userID, appID string) (*pbapplication.GetApplicationUserDetailResponse, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbapplication.NewApplicationManagementClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()
	resp, err := client.GetApplicationUserDetail(ctx, &pbapplication.GetApplicationUserDetailRequest{
		AppID:  appID,
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
