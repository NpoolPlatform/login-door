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

func Login(r *http.Request, request *mytype.LoginRequest) (*mytype.UserDetail, error) {
	resp, err := mygrpc.GetApplication(request.AppID)
	if err != nil {
		return nil, xerrors.Errorf("fail to get application info: %v", err)
	}

	if resp.Info.GoogleRecaptcha {
		if request.GoogleRecaptchaResponse == "" {
			return nil, xerrors.Errorf("didn't pass google recaptcha")
		}

		err := mygrpc.VerifyGoogleRecaptcha(request.GoogleRecaptchaResponse)
		if err != nil {
			return nil, err
		}
	}

	if request.Username != "" {
		return ByUsername(request)
	}
	if request.Email != "" {
		return ByEmail(request)
	}
	if request.Provider != "" {
		return ByThirdParty(request)
	}
	return nil, xerrors.Errorf("fail to login")
}

func ByUsername(request *mytype.LoginRequest) (*mytype.UserDetail, error) {
	resp, err := exist.User(request.Username, request.Password, request.AppID, "", "", false)
	if err != nil {
		return nil, err
	}

	err = mygrpc.VerifyCode(request.Username, request.VerifyCode)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ByEmail(request *mytype.LoginRequest) (*mytype.UserDetail, error) {
	resp, err := exist.User(request.Email, request.Password, request.AppID, "", "", false)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ByThirdParty(request *mytype.LoginRequest) (*mytype.UserDetail, error) {
	if request.Code == "" || request.State == "" {
		return nil, xerrors.Errorf("you need to auth in third party provider first")
	}

	providerInfo, err := provider.Get(context.Background(), request.Provider)
	if err != nil {
		return nil, err
	}

	idProvider := idp.GetIdProvider(providerInfo.ProviderName, providerInfo.ClientID, providerInfo.ClientSecret, request.RedirectURL)
	if idProvider == nil {
		return nil, xerrors.Errorf("get id provider empty")
	}

	token, err := idProvider.GetToken(request.Code)
	if err != nil {
		return nil, xerrors.Errorf("get provider token error: %v", err)
	}

	if !token.Valid() {
		return nil, xerrors.Errorf("invalid token")
	}

	providerUserInfo, err := idProvider.GetUserInfo(token)
	if err != nil {
		return nil, xerrors.Errorf("fail to login into third party provider: %v", err)
	}

	resp, err := exist.User("", "", request.AppID, request.Provider, providerUserInfo.Id, true)
	// provider user exist in our system.
	if err == nil && resp != nil {
		return resp, nil
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
	return nil, xerrors.Errorf("fail to login")
}

func GetUserLogin(request mytype.GetUserLoginRequest) (mytype.GetUserLoginResponse, error) {
	sessionContent, err := myredis.QueryKeyInfo(mytype.LoginKeyword, request.UserID[:8]+request.AppSession)
	if err == redis.Nil {
		return mytype.GetUserLoginResponse{}, nil
	}
	if err != nil {
		return mytype.GetUserLoginResponse{}, err
	}
	if sessionContent.UserID != request.UserID || sessionContent.Session != request.AppSession {
		return mytype.GetUserLoginResponse{}, xerrors.Errorf("user info not match")
	}

	return mytype.GetUserLoginResponse{
		Info: sessionContent,
	}, nil
}

func Logout(request mytype.LogoutRequest) error {
	resp, err := myredis.QueryKeyInfo(mytype.LoginKeyword, request.UserID[:8]+request.AppSession)
	if err != nil {
		return err
	}

	if resp.UserID != request.UserID {
		return xerrors.Errorf("user info doesn't match")
	}

	err = myredis.DelKey(mytype.LoginKeyword, request.UserID[:8]+request.AppSession)
	if err != nil {
		return err
	}

	return nil
}

func GetSSOLogin(request mytype.GetSSOLoginRequest) (mytype.GetSSOLoginResponse, error) {
	resp, err := myredis.QueryKeyInfo(mytype.LoginKeyword, request.UserID[:8]+request.Session)
	if err != nil {
		return mytype.GetSSOLoginResponse{}, err
	}

	if resp.Session != request.Session || resp.UserID != request.UserID {
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
