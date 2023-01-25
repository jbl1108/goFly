package driver

import (
	"time"

	"github.com/jbl1108/goFly/usecase"
	"github.com/jbl1108/goFly/util"
)

type newFetchFlightInfoAdapter struct {
	restClient       *RestClient
	mqttCommunicator *MQTTCommunicator
	config           *util.Config
}

func NewFetchFlightInfoAdapter(config *util.Config, restClient *RestClient) *newFetchFlightInfoAdapter {
	newFetchFlightInfoAdapter := new(newFetchFlightInfoAdapter)
	newFetchFlightInfoAdapter.config = config
	newFetchFlightInfoAdapter.restClient = restClient
	return newFetchFlightInfoAdapter
}

func (m *newFetchFlightInfoAdapter) Start() error {
	return m.mqttCommunicator.Start()
}

func (m *newFetchFlightInfoAdapter) PostMessage(message []usecase.FlightData) error {
	return nil
}

func (m *newFetchFlightInfoAdapter) SendFlightRequest(flightCode string, startDate time.Time, endDate time.Time) ([]usecase.FlightData, error) {
	parser := NewFlightDataParser()
	var request = m.config.FlightInfoRequest() + "/" + flightCode + "?start=" + startDate.Format(time.RFC3339) + "&end=" + endDate.Format(time.RFC3339)
	response, err := m.restClient.Request(request, map[string]string{"x-key": m.config.FlightInfoKey()})
	if err != nil {
		return nil, err
	} else {
		return parser.ParseData(response)
	}
}
