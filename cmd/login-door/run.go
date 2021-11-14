package main

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/NpoolPlatform/login-door/pkg/db"
	msgcli "github.com/NpoolPlatform/login-door/pkg/message/client"
	msglistener "github.com/NpoolPlatform/login-door/pkg/message/listener"
	msg "github.com/NpoolPlatform/login-door/pkg/message/message"
	msgsrv "github.com/NpoolPlatform/login-door/pkg/message/server"
	"github.com/go-chi/chi/v5"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/login-door/api"

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

		r := chi.NewMux()
		api.Register(r)
		err := http.ListenAndServe(fmt.Sprintf(":%v", HttpPort), r)
		if err != nil {
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

		return nil
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
