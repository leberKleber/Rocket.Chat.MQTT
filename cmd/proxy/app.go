package main

import (
	"context"
	"encoding/json"
	"github.com/leberKleber/Rocket.Chat.MQTT/internal/mqtt"
	"github.com/leberKleber/Rocket.Chat.MQTT/internal/rocketchat"
	"github.com/leberKleber/Rocket.Chat.MQTT/internal/rocketchat/message"
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

	mqttClient := mqtt.NewClient(cfg.BrokerURL, cfg.ClientID)
	defer shutdown(rcClient, mqttClient)

	err = rcClient.Start()
	if err != nil {
		log.WithError(err).Fatal("Failed to start rocketChat client")
		os.Exit(1)
		return
	}

	err = mqttClient.Start()
	if err != nil {
		log.WithError(err).Fatal("Failed to start mqtt client")
		os.Exit(1)
		return
	}

	err = rcClient.Login(cfg.Username, cfg.PasswordHash)
	if err != nil {
		log.WithError(err).Fatal("Failed to login")
		os.Exit(1)
		return
	}

	msg := message.NewGetRooms()
	resp := rcClient.SendMessageWaitForResponse(msg.ID, msg)

	rresp := message.GetRoomsResponse{}

	err = json.Unmarshal(resp, &rresp)
	log.Info(err)

	var rooms = make(map[string]string)

	for _, r := range rresp.Results {
		if r.Type == "c" {
			rooms[r.Name] = r.ID
		}
	}

	rcClient.SendMessage(message.NewSendMessage(rooms["testav"]))

	if resp != nil {
		log.Info(string(resp))
	} else {
		log.Info("NIL response")
	}

	for {
		time.Sleep(time.Minute * 5)
	}
}

func shutdown(rcClient rocketchat.Client, mqttClient mqtt.MqttClient) {
	err := rcClient.Stop()
	if err != nil {
		log.WithError(err).Error("Error while stopping rocketchat client")
	}

	err = mqttClient.Stop()
	if err != nil {
		log.WithError(err).Error("Error while stopping mqtt client")
	}
}
