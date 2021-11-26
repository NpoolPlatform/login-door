package exist

import (
	mygrpc "github.com/NpoolPlatform/login-door/pkg/grpc"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	pbuser "github.com/NpoolPlatform/user-management/message/npool"
	"golang.org/x/xerrors"
)

func User(username, password, appID, providerID, providerUserID string, thirdParty bool) (*mytype.UserDetail, error) {
	userBasicInfo := &pbuser.UserBasicInfo{} // nolint
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
		userBasicInfo = resp
	} else {
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

		userBasicInfo = resp
	}

	resp, err := mygrpc.GetUserDetail(userBasicInfo.UserID, appID)
	if err != nil {
		return nil, err
	}

	return &mytype.UserDetail{
		UserBasicInfo: userBasicInfo,
		UserAppInfo:   resp.Info,
	}, nil
}
