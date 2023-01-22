package driver

import (
	"github.com/jbl1108/goFly/usecase"
)

type newFetchFlightInfoAdapter struct {
	restClient       restClient
	mqttCommunicator MQTTCommunicator
	aviationHost     string
}

func NewFetchFlightInfoAdapter(aviationHost string, mqqttHost string) *newFetchFlightInfoAdapter {
	newFetchFlightInfoAdapter := new(newFetchFlightInfoAdapter)
	newFetchFlightInfoAdapter.restClient = *NewRestClient()
	newFetchFlightInfoAdapter.aviationHost = aviationHost
	newFetchFlightInfoAdapter.mqttCommunicator = *NewMQTTCommunicator(mqqttHost, "flyinfo")
	return newFetchFlightInfoAdapter
}

func (m *newFetchFlightInfoAdapter) Start() error {
	return m.mqttCommunicator.Start()
}
func (m *newFetchFlightInfoAdapter) RegisterMessageListener(topic string, listener func(message string)) {

}
func (m *newFetchFlightInfoAdapter) PostMessage(message string, receiver string) error {
	return nil
}

func (m *newFetchFlightInfoAdapter) SendRequest(message string) ([]usecase.FlightData, error) {
	parser := NewFlightDataParser()
	response, err := m.restClient.Request(m.aviationHost)
	if err == nil {
		return nil, err
	} else {
		return parser.ParseData(response)
	}
}
