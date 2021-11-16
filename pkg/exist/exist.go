package exist

import (
	mygrpc "github.com/NpoolPlatform/login-door/pkg/grpc"
	"golang.org/x/xerrors"
)

func User(username, password, appID, providerID, providerUserID string, thirdParty bool) (string, error) {
	var userID string
	if !thirdParty {
		resp, err := mygrpc.QueryUserExist(username, password)
		if err != nil {
			return "", xerrors.Errorf("query user exist error: %v", err)
		}

		err = mygrpc.QueryUserInApplication(resp.Info.UserId, appID)
		if err != nil {
			return "", xerrors.Errorf("user can not login into app: %v", err)
		}

		err = mygrpc.QueryUserFrozen(resp.Info.UserId)
		if err != nil {
			return "", err
		}
		userID = resp.Info.UserId
	} else {
		resp, err := mygrpc.QueryUserByUserProviderID(providerID, providerUserID)
		if err != nil {
			return "", nil
		}

		err = mygrpc.QueryUserInApplication(resp.Info.UserProviderInfo.UserId, appID)
		if err != nil {
			return "", xerrors.Errorf("user can not login into app: %v", err)
		}

		err = mygrpc.QueryUserFrozen(resp.Info.UserProviderInfo.UserId)
		if err != nil {
			return "", err
		}
		userID = resp.Info.UserProviderInfo.UserId
	}
	return userID, nil
}