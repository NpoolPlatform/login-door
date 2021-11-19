package login

import (
	"context"
	"net/http"

	"github.com/NpoolPlatform/login-door/pkg/crud/provider"
	"github.com/NpoolPlatform/login-door/pkg/exist"
	mygrpc "github.com/NpoolPlatform/login-door/pkg/grpc"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	myredis "github.com/NpoolPlatform/login-door/pkg/redis"
	"github.com/casbin/casdoor/idp"
	"github.com/go-redis/redis/v8"
	"golang.org/x/xerrors"
)

func Login(r *http.Request, request *mytype.LoginRequest) (string, error) {
	if request.Username != "" {
		return ByUsername(request)
	}
	if request.Email != "" {
		return ByEmailVerifyCode(request)
	}
	if request.Provider != "" {
		return ByThirdParty(request)
	}
	return "", xerrors.Errorf("fail to login")
}

func ByUsername(request *mytype.LoginRequest) (string, error) {
	userID, err := exist.User(request.Username, request.Password, request.AppID, "", "", false)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func ByEmailVerifyCode(request *mytype.LoginRequest) (string, error) {
	if request.VerifyCode == "" {
		return "", xerrors.Errorf("verify code can not be empty")
	}

	err := mygrpc.VerifyCode(request.Email, request.VerifyCode)
	if err != nil {
		return "", xerrors.Errorf("verify code is wrong")
	}

	userID, err := exist.User(request.Email, "", request.AppID, "", "", false)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func ByThirdParty(request *mytype.LoginRequest) (string, error) {
	if request.Code == "" || request.State == "" {
		return "", xerrors.Errorf("you need to auth in third party provider first")
	}

	providerInfo, err := provider.Get(context.Background(), request.Provider)
	if err != nil {
		return "", err
	}

	idProvider := idp.GetIdProvider(providerInfo.ProviderName, providerInfo.ClientID, providerInfo.ClientSecret, request.RedirectURL)
	if idProvider == nil {
		return "", xerrors.Errorf("get id provider empty")
	}

	token, err := idProvider.GetToken(request.Code)
	if err != nil {
		return "", xerrors.Errorf("get provider token error: %v", err)
	}

	if !token.Valid() {
		return "", xerrors.Errorf("invalid token")
	}

	providerUserInfo, err := idProvider.GetUserInfo(token)
	if err != nil {
		return "", xerrors.Errorf("fail to login into third party provider: %v", err)
	}

	userID, err := exist.User("", "", request.AppID, request.Provider, providerUserInfo.Id, true)
	// provider user exist in our system.
	if err == nil && userID != "" {
		return userID, nil
	}
	// provider user id doesn't exist in our system.
	// provider has still not bind to a user.
	// TODO: create a user by user email.
	// if err == nil && userID == "" {
	// 	resp, errEmail := mygrpc.QueryUserExist(providerUserInfo.Email, "")
	// 	if err != nil {
	// 		// provider user doesn't exist by query its email.
	// 	}
	// }
	return "", xerrors.Errorf("fail to login")
}

func GetUserLogin(request mytype.GetUserLoginRequest) (mytype.GetUserLoginResponse, error) {
	sessionContent, err := myredis.QueryKeyInfo(mytype.LoginKeyword, request.AppLoginSession)
	if err == redis.Nil {
		return mytype.GetUserLoginResponse{}, nil
	}
	if err != nil {
		return mytype.GetUserLoginResponse{}, err
	}
	if sessionContent.UserID != request.UserID || sessionContent.Session != request.AppLoginSession {
		return mytype.GetUserLoginResponse{}, xerrors.Errorf("user info not match")
	}

	return mytype.GetUserLoginResponse{
		Info: sessionContent,
	}, nil
}

func Logout(request mytype.LogoutRequest) error {
	resp, err := myredis.QueryKeyInfo(mytype.LoginKeyword, request.Session)
	if err != nil {
		return err
	}

	if resp.UserID != request.UserID {
		return xerrors.Errorf("user info doesn't match")
	}

	err = myredis.DelKey(mytype.LoginKeyword, request.Session)
	if err != nil {
		return err
	}

	return nil
}

func GetSSOLogin(request mytype.GetSSOLoginRequest) (mytype.GetSSOLoginResponse, error) {
	resp, err := myredis.QueryKeyInfo(mytype.LoginKeyword, request.LoginSession)
	if err != nil {
		return mytype.GetSSOLoginResponse{}, err
	}

	if resp.Session != request.LoginSession || resp.UserID != request.UserID {
		return mytype.GetSSOLoginResponse{}, xerrors.Errorf("invalid user")
	}

	err = mygrpc.QueryUserInApplication(request.UserID, request.AppID)
	if err != nil {
		return mytype.GetSSOLoginResponse{}, xerrors.Errorf("user can not login into app: %v", err)
	}

	return mytype.GetSSOLoginResponse{
		Info: resp,
	}, nil
}
