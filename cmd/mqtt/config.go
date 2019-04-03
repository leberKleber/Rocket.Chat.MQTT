package main

import "github.com/alexflint/go-arg"

type Config struct {
	RocketChat
	MQTT
}

type RocketChat struct {
	WsURL        string `arg:"env:ROCKET_CHAT_WS_URL"`
	Username     string `arg:"env:ROCKET_CHAT_USERNAME"`
	PasswordHash string `arg:"env:ROCKET_CHAT_PASSWORD_HASH"`
}

type MQTT struct {
	BrokerURL string `arg:"env:MQTT_BROKER_URL"`
	ClientID  string `arg:"env:MQTT_CLIENT_ID"`
}

func NewConfig() (Config, error) {
	cfg := Config{}
	err := arg.Parse(&cfg)

	return cfg, err
}
