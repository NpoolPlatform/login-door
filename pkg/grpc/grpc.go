package grpc

import (
	"context"
	"encoding/json"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pbUser "github.com/NpoolPlatform/user-management/message/npool"
	userconst "github.com/NpoolPlatform/user-management/pkg/message/const"
	pbVerification "github.com/NpoolPlatform/verification-door/message/npool"
	verificationconst "github.com/NpoolPlatform/verification-door/pkg/message/const"
	"github.com/casbin/casdoor/idp"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

const (
	VerificationService     = verificationconst.ServiceName
	VerificationServicePort = "50091"
	UserService             = userconst.ServiceName
	UserServicePort         = "50071"
	ApplicationService      = applicationconst.ServiceName
	ApplicationServicePort  = "50081"
)

func newVerificationGrpcConn() (*grpc.ClientConn, error) {
	conn, err := mygrpc.GetGRPCConn(verificationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func VerifyCode(param, code string) error {
	conn, err := newVerificationGrpcConn()
	if err != nil {
		return err
	}

	client := pbVerification.NewVerificationDoorClient(conn)
	_, err = client.VerifyCode(context.Background(), &pbVerification.VerifyCodeRequest{
		Param: param,
		Code:  code,
	})
	if err != nil {
		return err
	}
	return nil
}

func newUserGrpcConn() (*grpc.ClientConn, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CreateTestUser(appID string) (*pbUser.SignupResponse, error) {
	conn, err := newUserGrpcConn()
	if err != nil {
		return nil, err
	}

	client := pbUser.NewUserClient(conn)

	resp, err := client.SignUp(context.Background(), &pbUser.SignupRequest{
		AppId:    appID,
		Username: uuid.New().String(),
		Password: "12345679",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CreateUser(appID, providerID string, providerUserInfo *idp.UserInfo) (*pbUser.BindThirdPartyResponse, error) {
	conn, err := newUserGrpcConn()
	if err != nil {
		return nil, err
	}

	client := pbUser.NewUserClient(conn)

	resp, err := client.AddUser(context.Background(), &pbUser.AddUserRequest{
		AppId: appID,
		UserInfo: &pbUser.UserBasicInfo{
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
	respBind, err := client.BindThirdParty(context.Background(), &pbUser.BindThirdPartyRequest{
		UserId:           resp.Info.UserId,
		AppId:            appID,
		ProviderId:       providerID,
		ProviderUserId:   providerUserInfo.Id,
		UserProviderInfo: string(b),
	})
	if err != nil {
		return nil, err
	}

	return respBind, nil
}

func QueryUserExist(username, password string) (*pbUser.QueryUserExistResponse, error) {
	conn, err := newUserGrpcConn()
	if err != nil {
		return nil, err
	}

	client := pbUser.NewUserClient(conn)
	resp, err := client.QueryUserExist(context.Background(), &pbUser.QueryUserExistRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func QueryUserByUserProviderID(providerID, userProviderID string) (*pbUser.QueryUserByUserProviderIDResponse, error) {
	conn, err := newUserGrpcConn()
	if err != nil {
		return nil, err
	}

	client := pbUser.NewUserClient(conn)
	resp, err := client.QueryUserByUserProviderID(context.Background(), &pbUser.QueryUserByUserProviderIDRequest{
		ProviderID:     providerID,
		ProviderUserID: userProviderID,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func QueryUserFrozen(userID string) error {
	conn, err := newUserGrpcConn()
	if err != nil {
		return err
	}

	client := pbUser.NewUserClient(conn)
	resp, err := client.QueryUserFrozen(context.Background(), &pbUser.QueryUserFrozenRequest{
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

func newApplicationGrpcConn() (*grpc.ClientConn, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CreaeteApp() (*pbApplication.CreateApplicationResponse, error) {
	conn, err := newApplicationGrpcConn()
	if err != nil {
		return nil, err
	}

	client := pbApplication.NewApplicationManagementClient(conn)
	resp, err := client.CreateApplication(context.Background(), &pbApplication.CreateApplicationRequest{
		Info: &pbApplication.ApplicationInfo{
			ApplicationName:  uuid.New().String(),
			ApplicationOwner: uuid.New().String(),
		},
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func QueryUserInApplication(userID, appID string) error {
	conn, err := newApplicationGrpcConn()
	if err != nil {
		return err
	}

	client := pbApplication.NewApplicationManagementClient(conn)
	_, err = client.GetUserFromApplication(context.Background(), &pbApplication.GetUserFromApplicationRequest{
		AppID:  appID,
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return nil
}
