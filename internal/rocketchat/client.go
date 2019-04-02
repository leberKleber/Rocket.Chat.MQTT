package rocketchat

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/leberKleber/Rocket.Chat.MQTT/internal/rocketchat/message"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var incomingMessages = make(map[string][]byte)

type wsMessage struct {
	MessageType    int
	MessagePayload []byte
}

type Client struct {
	ctx              context.Context
	wsConnection     *websocket.Conn
	incomingMessages chan wsMessage
}

func NewClient(ctx context.Context, url string) (Client, error) {
	wsConnection, httpResponse, err := websocket.DefaultDialer.Dial(url, nil)
	if httpResponse == nil || httpResponse.StatusCode != http.StatusSwitchingProtocols {
		return Client{}, errors.New("Could not establish websocket connection")
	}

	if err != nil {
		return Client{}, errors.Wrap(err, "Could not establish websocket connection")
	}

	return Client{
		ctx:          ctx,
		wsConnection: wsConnection,
	}, nil
}

func (rcc *Client) handleIncoming(msgType int, msgPayload []byte) {
	switch msgType {
	case websocket.TextMessage:
		var msg message.GeneralMessage
		err := json.Unmarshal(msgPayload, &msg)
		if err != nil {
			log.WithError(err).Error("Failed to unmarshal msg payload")
		}

		log.Infof("Incoming text-message: %d  %v", msg, string(msgPayload))

		switch msg.Message {
		case "ping":
			log.Debug("Send WS pong")
			err = rcc.wsConnection.WriteJSON(message.NewPong()) //RocketChat has a Ping/Ping/Pong flow (Server>Client/Client>Server/Server>Client)
			if err != nil {
				log.WithError(err).Error("Failed to send pong")
			}
		default:
			incomingMessages[msg.ID] = msgPayload
		}

		break
	case websocket.BinaryMessage:
		log.Info("Incoming binary-message")
		break
	case websocket.PingMessage:
		log.Info("Incoming ping-message")
		break
	default:
		log.Warnf("Unexpected message-type received: %d", msgType)
	}
}

func (rcc *Client) Start() error {
	go func() {
		type response struct {
			messageType int
			message     []byte
			err         error
		}
		internalChan := make(chan response)
		defer close(internalChan)

		for {
			//wrapper go func to use the channel functionality
			go func() {
				resp := response{}
				resp.messageType, resp.message, resp.err = rcc.wsConnection.ReadMessage()
				internalChan <- resp
			}()

			select {
			case <-rcc.ctx.Done():
				//close signal received: that means we have to go
				return
			case resp := <-internalChan:
				if resp.err == nil {
					rcc.handleIncoming(resp.messageType, resp.message)
				}
			}
		}
	}()

	return rcc.wsConnection.WriteJSON(message.NewConnect())
}

func (rcc *Client) SendMessage(msg interface{}) error {
	return rcc.wsConnection.WriteJSON(msg)
}

func (rcc *Client) SendMessageWaitForResponse(msgID string, msg interface{}) []byte {
	rcc.SendMessage(msg)
	for {
		resp := incomingMessages[msgID]
		if resp == nil {
			time.Sleep((1 * time.Second) / 2)
		} else {
			return resp
		}
	}
}

func (rcc *Client) Stop() error {
	if rcc.wsConnection != nil {
		return rcc.wsConnection.Close()
	}

	return nil
}

func (rcc *Client) Login(username, passwordHash string) error {
	return rcc.wsConnection.WriteJSON(message.NewLogin(username, passwordHash))
}
