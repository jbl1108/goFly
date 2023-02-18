package delivery

import (
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MQTTCommunicator struct {
	messageListener   func(message string)
	brokerHostAddress string
}

func NewMQTTCommunicator(brokerHostAddress string) *MQTTCommunicator {
	mqttcommunicator := new(MQTTCommunicator)
	mqttcommunicator.brokerHostAddress = brokerHostAddress
	return mqttcommunicator
}

var mqttClient MQTT.Client

func (m *MQTTCommunicator) RegisterListener(topic string, listener func(message string)) {
	m.messageListener = listener
}
func (m *MQTTCommunicator) SendMessage(message string, receiver string) error {
	token := mqttClient.Publish(receiver, 0, false, message)
	token.Wait()
	return token.Error()
}

func (m *MQTTCommunicator) Start() error {
	opts := MQTT.NewClientOptions().AddBroker(m.brokerHostAddress)
	opts.SetClientID("goFly client")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		log.Printf("[CONNECTION LOST HANDLER] %v", err)
	})

	var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		m.messageListener(string(msg.Payload()))
	}
	opts.SetDefaultPublishHandler(f)
	//create and start a client using the above ClientOptions
	mqttClient = MQTT.NewClient(opts)

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	//subscribe to the topic catdata and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	// if token := mqttClient.Subscribe(m.topic, 0, nil); token.Wait() && token.Error() != nil {
	// 	return token.Error()
	// }
	return nil
}
func (m *MQTTCommunicator) Stop() {
	// if token := mqttClient.Unsubscribe(m.topic); token.Wait() && token.Error() != nil {
	// 	log.Printf("Error disconnecting: %v", token.Error())
	// }

	mqttClient.Disconnect(250)

}
