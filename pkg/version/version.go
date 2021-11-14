package version

import (
	"github.com/NpoolPlatform/login-door/pkg/mytype"

	cv "github.com/NpoolPlatform/go-service-framework/pkg/version"

	"golang.org/x/xerrors"
)

func Version() (mytype.VersionResponse, error) {
	info, err := cv.GetVersion()
	if err != nil {
		return mytype.VersionResponse{}, xerrors.Errorf("get service version error: %w", err)
	}
	return mytype.VersionResponse{
		Info: info,
	}, nil
}
