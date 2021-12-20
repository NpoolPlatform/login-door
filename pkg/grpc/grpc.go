package grpc

import (
	"context"
	"encoding/json"

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

func VerifyCode(param, code string) error {
	conn, err := mygrpc.GetGRPCConn(verificationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbverification.NewVerificationDoorClient(conn)
	_, err = client.VerifyCode(context.Background(), &pbverification.VerifyCodeRequest{
		Param: param,
		Code:  code,
	})
	if err != nil {
		return err
	}
	return nil
}

func VerifyGoogleRecaptcha(response string) error {
	conn, err := mygrpc.GetGRPCConn(verificationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbverification.NewVerificationDoorClient(conn)
	resp, err := client.VerifyGoogleRecaptcha(context.Background(), &pbverification.VerifyGoogleRecaptchaRequest{
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

func CreateTestUser(appID string) (*pbuser.SignupResponse, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)
	resp, err := client.SignUp(context.Background(), &pbuser.SignupRequest{
		AppID:    appID,
		Username: uuid.New().String(),
		Password: "12345679",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CreateUser(appID, providerID string, providerUserInfo *idp.UserInfo) (*pbuser.BindThirdPartyResponse, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)
	resp, err := client.AddUser(context.Background(), &pbuser.AddUserRequest{
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
	respBind, err := client.BindThirdParty(context.Background(), &pbuser.BindThirdPartyRequest{
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

func QueryUserExist(username, password string) (*pbuser.UserBasicInfo, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)
	resp, err := client.QueryUserExist(context.Background(), &pbuser.QueryUserExistRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return resp.Info, nil
}

func QueryUserByUserProviderID(providerID, userProviderID string) (*pbuser.UserBasicInfo, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)
	resp, err := client.QueryUserByUserProviderID(context.Background(), &pbuser.QueryUserByUserProviderIDRequest{
		ProviderID:     providerID,
		ProviderUserID: userProviderID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Info.UserBasicInfo, nil
}

func QueryUserFrozen(userID string) error {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)
	resp, err := client.QueryUserFrozen(context.Background(), &pbuser.QueryUserFrozenRequest{
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

func CreaeteApp() (*pbapplication.CreateApplicationResponse, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbapplication.NewApplicationManagementClient(conn)
	resp, err := client.CreateApplication(context.Background(), &pbapplication.CreateApplicationRequest{
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

func GetApplication(appID string) (*pbapplication.GetApplicationResponse, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbapplication.NewApplicationManagementClient(conn)
	resp, err := client.GetApplication(context.Background(), &pbapplication.GetApplicationRequest{
		AppID: appID,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func QueryUserInApplication(userID, appID string) error {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbapplication.NewApplicationManagementClient(conn)
	_, err = client.GetUserFromApplication(context.Background(), &pbapplication.GetUserFromApplicationRequest{
		AppID:  appID,
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserDetail(userID, appID string) (*pbapplication.GetApplicationUserDetailResponse, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbapplication.NewApplicationManagementClient(conn)
	resp, err := client.GetApplicationUserDetail(context.Background(), &pbapplication.GetApplicationUserDetailRequest{
		AppID:  appID,
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
