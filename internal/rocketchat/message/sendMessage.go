package message

import (
	"github.com/google/uuid"
)

/*
{
    "msg": "method",
    "method": "sendMessage",
    "id": "42",
    "params": [
        {
            "_id": "message-id",
            "rid": "room-id",
            "msg": "Hello World!"
        }
    ]
}
*/
type SendMessage struct {
	baseMethodMessage
	Parameters []SendMessageParameters `json:"params"`
}

type SendMessageParameters struct {
	ID      string `json:"_id"`
	RoomID  string `json:"rid"`
	Message string `json:"msg"`
}

func NewSendMessage(roomID string) SendMessage {
	return SendMessage{
		baseMethodMessage: baseMethodMessage{
			Message: "method",
			Method:  "sendMessage",
			ID:      uuid.New().String(),
		},
		Parameters: []SendMessageParameters{
			{
				Message: "TestMessage32623462436236",
				RoomID:  roomID,
				ID:      uuid.New().String(),
			},
		},
	}
}
