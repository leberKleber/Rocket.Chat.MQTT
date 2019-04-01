package message

import (
	"github.com/google/uuid"
)

type Login struct {
	baseMethodMessage
	Parameters []LoginParameter `json:"params"`
}

type LoginParameter struct {
	User     LoginParameterUser     `json:"user"`
	Password LoginParameterPassword `json:"password"`
}

type LoginParameterUser struct {
	Username string `json:"username"`
}

type LoginParameterPassword struct {
	Digest    string `json:"digest"`
	Algorithm string `json:"algorithm"`
}

func NewLogin(username, passwordHash string) Login {
	return Login{
		baseMethodMessage: baseMethodMessage{
			Message: "method",
			Method:  "login",
			ID:      uuid.New().String(),
		},
		Parameters: []LoginParameter{
			{
				User: LoginParameterUser{
					Username: username,
				},
				Password: LoginParameterPassword{
					Algorithm: "sha-256",
					Digest:    passwordHash,
				},
			},
		},
	}
}
