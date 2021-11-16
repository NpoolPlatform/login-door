package testinit

import (
	"fmt"
	"path"
	"runtime"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/login-door/pkg/db"
	servicename "github.com/NpoolPlatform/login-door/pkg/service-name"
	"golang.org/x/xerrors"
)

func Init() error {
	_, myPath, _, ok := runtime.Caller(0)
	if !ok {
		return xerrors.Errorf("cannot get source file path")
	}

	appName := path.Base(path.Dir(path.Dir(path.Dir(myPath))))
	configPath := fmt.Sprintf("%s/../../cmd/%v", path.Dir(myPath), appName)
	fmt.Println("appname is:", appName, "config path is", configPath)

	err := app.Init(servicename.ServiceName, "", "", "", configPath, nil, nil)
	if err != nil {
		return xerrors.Errorf("cannot init app stub: %v", err)
	}
	err = db.Init()
	if err != nil {
		return xerrors.Errorf("cannot init database: %v", err)
	}

	return nil
}