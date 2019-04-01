package main

import "github.com/alexflint/go-arg"

type Config struct {
	RocketChat
}

type RocketChat struct {
	WsURL        string `arg:"env:ROCKET_CHAT_WS_URL"`
	Username     string `arg:"env:ROCKET_CHAT_USERNAME"`
	PasswordHash string `arg:"env:ROCKET_CHAT_PASSWORD_HASH"`
}

func NewConfig() (Config, error) {
	cfg := Config{}
	err := arg.Parse(&cfg)

	return cfg, err
}
