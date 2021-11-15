package grpc

import (
	"context"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pbPermission "github.com/NpoolPlatform/permission-door/message/npool"
	pbUser "github.com/NpoolPlatform/user-management/message/npool"
	pbVerification "github.com/NpoolPlatform/verification-door/message/npool"
	"google.golang.org/grpc"
)

const (
	VerificationService     = "verification-door.npool.top"
	VerificationServicePort = ":50091"
	PermissionService       = "permission-door.npool.top"
	PermissionServicePort   = ":50101"
	UserService             = "user-management.npool.top"
	UserServicePort         = ":50071"
)

func newVerificationGrpcConn() (*grpc.ClientConn, error) {
	serviceAgent, err := config.PeekService(VerificationService)
	if err != nil {
		return nil, err
	}

	myAddress := []string{}
	for _, address := range strings.Split(serviceAgent.Address, ",") {
		myAddress = append(myAddress, address+VerificationServicePort)
	}

	conn, err := mygrpc.GetGRPCConn(strings.Join(myAddress, ","))
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

	defer conn.Close()

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

func newPermissionGrpcConn() (*grpc.ClientConn, error) {
	serviceAgent, err := config.PeekService(PermissionService)
	if err != nil {
		return nil, err
	}

	myAddress := []string{}
	for _, address := range strings.Split(serviceAgent.Address, ",") {
		myAddress = append(myAddress, address+PermissionServicePort)
	}

	conn, err := mygrpc.GetGRPCConn(strings.Join(myAddress, ","))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func AuthenticateUserByUserID(userID, appID, resourceID, action string) error {
	conn, err := newPermissionGrpcConn()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbPermission.NewPermissionDoorClient(conn)
	_, err = client.AuthenticateUserPolicyByID(context.Background(), &pbPermission.AuthenticateUserPolicyByIDRequest{
		UserID:     userID,
		AppID:      appID,
		ResourecID: resourceID,
		Action:     action,
	})
	if err != nil {
		return err
	}
	return nil
}

func newUserGrpcConn() (*grpc.ClientConn, error) {
	serviceAgent, err := config.PeekService(UserService)
	if err != nil {
		return nil, err
	}

	myAddress := []string{}
	for _, address := range strings.Split(serviceAgent.Address, ",") {
		myAddress = append(myAddress, address+UserServicePort)
	}

	conn, err := mygrpc.GetGRPCConn(strings.Join(myAddress, ","))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func QueryUserExist(username, password string) (*pbUser.QueryUserExistResponse, error) {
	conn, err := newUserGrpcConn()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

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

	defer conn.Close()

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
