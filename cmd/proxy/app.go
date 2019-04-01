package main

import (
	"context"
	"github.com/leberKleber/Rocket.Chat.MQTT/internal/rocketchat"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.WithError(err).Fatal("Failed to fetch config")
		os.Exit(1)
		return
	}

	rcClient, err := rocketchat.NewClient(context.Background(), cfg.WsURL)
	if err != nil {
		log.WithError(err).Fatal("Failed to create rocketChat client")
		os.Exit(1)
		return
	}

	err = rcClient.Start()
	if err != nil {
		log.WithError(err).Fatal("Failed to start rocketChat client")
		os.Exit(1)
		return
	}

	err = rcClient.Login(cfg.Username, cfg.PasswordHash)
	if err != nil {
		log.WithError(err).Fatal("Failed to login")
		os.Exit(1)
		return
	}

	for {
		time.Sleep(time.Minute * 5)
	}
}
