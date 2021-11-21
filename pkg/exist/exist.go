package exist

import (
	mygrpc "github.com/NpoolPlatform/login-door/pkg/grpc"
	pbUser "github.com/NpoolPlatform/user-management/message/npool"
	"golang.org/x/xerrors"
)

func User(username, password, appID, providerID, providerUserID string, thirdParty bool) (*pbUser.UserBasicInfo, error) {
	if !thirdParty {
		resp, err := mygrpc.QueryUserExist(username, password)
		if err != nil {
			return nil, xerrors.Errorf("query user exist error: %v", err)
		}

		err = mygrpc.QueryUserInApplication(resp.UserID, appID)
		if err != nil {
			return nil, xerrors.Errorf("user can not login into app: %v", err)
		}

		err = mygrpc.QueryUserFrozen(resp.UserID)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	resp, err := mygrpc.QueryUserByUserProviderID(providerID, providerUserID)
	if err != nil {
		return nil, nil
	}

	err = mygrpc.QueryUserInApplication(resp.UserID, appID)
	if err != nil {
		return nil, xerrors.Errorf("user can not login into app: %v", err)
	}

	err = mygrpc.QueryUserFrozen(resp.UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
