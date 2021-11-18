package main

import (
	"time"

	db "github.com/NpoolPlatform/login-door/pkg/db"
	msgcli "github.com/NpoolPlatform/login-door/pkg/message/client"
	msglistener "github.com/NpoolPlatform/login-door/pkg/message/listener"
	msg "github.com/NpoolPlatform/login-door/pkg/message/message"
	msgsrv "github.com/NpoolPlatform/login-door/pkg/message/server"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/login-door/api"

	http2 "github.com/NpoolPlatform/go-service-framework/pkg/http"
	cli "github.com/urfave/cli/v2"
)

const HttpPort = 50060

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		if err := db.Init(); err != nil {
			return err
		}

		if err := msgsrv.Init(); err != nil {
			return err
		}
		if err := msgcli.Init(); err != nil {
			return err
		}

		go msglistener.Listen()
		go msgSender()

		return http2.Run(api.Register)
	},
}

func msgSender() {
	id := 0
	for {
		logger.Sugar().Infof("send example")
		err := msgsrv.PublishExample(&msg.Example{
			ID:      id,
			Example: "hello world",
		})
		if err != nil {
			logger.Sugar().Errorf("fail to send example: %v", err)
			return
		}
		id++
		time.Sleep(3 * time.Second)
	}
}
