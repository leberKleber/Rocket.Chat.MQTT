package mqtt

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

type MqttClient struct {
	mqttTopicPrefix    string
	cli                *client.Client
	mqttConnectOptions *client.ConnectOptions
}

func NewClient(mqttBroker, clientID string) MqttClient {
	mc := MqttClient{}
	cli := client.New(&client.Options{
		ErrorHandler: func(err error) {
			mc.handleMqttError(err)
		},
	})

	mc.cli = cli
	mc.mqttConnectOptions = &client.ConnectOptions{
		Network:  "tcp",
		Address:  mqttBroker,
		ClientID: []byte(clientID),
	}
	return mc
}

func (mc *MqttClient) Start() error {
	return mc.cli.Connect(mc.mqttConnectOptions)
}

func (mc *MqttClient) Stop() error {
	err := mc.cli.Disconnect()
	mc.cli.Terminate()

	return err
}

func (mc *MqttClient) PublishJSON(relativeTopic string, payload interface{}) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		log.WithError(err).Error("Failed to marshal json to publish")
		return
	}

	err = mc.cli.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS1,
		TopicName: []byte(mc.mqttTopicPrefix + relativeTopic),
		Message:   bytes,
	})
	if err != nil {
		log.WithError(err).Error("Failed to publish message")
		return
	}
}

func (mc *MqttClient) Subscribe(relativeTopic string, handler client.MessageHandler) error {
	return mc.cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			{
				QoS:         1,
				TopicFilter: []byte(relativeTopic),
				Handler:     handler,
			},
		},
	})
}

func (mc *MqttClient) handleMqttError(err error) {
	log.WithError(err).Error("An error has occurred")
}
