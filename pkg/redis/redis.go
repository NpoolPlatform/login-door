package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/login-door/pkg/mytype"
	"github.com/go-redis/redis/v8"
	"golang.org/x/xerrors"
)

func Client() *redis.Client {
	return app.Redis().Client
}

func InsertKeyInfo(keyword, param string, info interface{}, ttl time.Duration) error {
	b, err := json.Marshal(info)
	if err != nil {
		fmt.Println("json error is", err)
		return err
	}
	err = Client().Set(context.Background(), fmt.Sprintf("%v::%v", keyword, param), string(b), ttl).Err()
	if err != nil {
		fmt.Println("set error is:", err)
		return err
	}
	return nil
}

func QueryKeyInfo(keyword, param string) (mytype.LoginSession, error) {
	val, err := Client().Get(context.Background(), fmt.Sprintf("%v::%v", keyword, param)).Result()
	if err != nil {
		return mytype.LoginSession{}, xerrors.Errorf("fail to get redis key: %v", err)
	}

	response := mytype.LoginSession{}
	err = json.Unmarshal([]byte(val), &response)

	if err == redis.Nil {
		return mytype.LoginSession{}, nil
	} else if err != nil {
		return mytype.LoginSession{}, xerrors.Errorf("fail to unmarshal json: %v", err)
	}

	return response, nil
}

func DelKey(keyword, param string) error {
	err := Client().Del(context.Background(), fmt.Sprintf("%v::%v", keyword, param)).Err()
	if err != nil {
		return xerrors.Errorf("delete key error: %v", err)
	}
	return nil
}
