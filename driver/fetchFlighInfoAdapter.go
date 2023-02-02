package driver

import (
	"fmt"
	"log"
	"time"

	"github.com/jbl1108/goFly/usecase"
	"github.com/jbl1108/goFly/util"
	"go.uber.org/multierr"
)

type newFetchFlightInfoAdapter struct {
	restClient       *RestClient
	mqttCommunicator *MQTTCommunicator
	config           *util.Config
}

func NewFetchFlightInfoAdapter(config *util.Config, restClient *RestClient, mqttClient *MQTTCommunicator) *newFetchFlightInfoAdapter {
	newFetchFlightInfoAdapter := new(newFetchFlightInfoAdapter)
	newFetchFlightInfoAdapter.config = config
	newFetchFlightInfoAdapter.restClient = restClient
	newFetchFlightInfoAdapter.mqttCommunicator = mqttClient
	return newFetchFlightInfoAdapter
}

func (m *newFetchFlightInfoAdapter) Start() error {
	return m.mqttCommunicator.Start()
}

func (m *newFetchFlightInfoAdapter) PostMessage(message []usecase.FlightData) error {
	var errors error
	for _, flight := range message {
		json := m.generateJson(flight)
		topic := "flight/" + flight.IataFlightCode
		log.Printf("Post topic: %s, message : %s", topic, json)
		err := m.mqttCommunicator.SendMessage(json, topic)
		errors = multierr.Append(errors, err)
	}
	return errors
}

func (m *newFetchFlightInfoAdapter) generateJson(flight usecase.FlightData) string {
	return fmt.Sprintf("{\"time\" : \"%s\" , \"text_flightDate\" : \"%s\" , \"arrivalDelay\" : %f , \"departureDelay\" : %f }", flight.FlightDate.Format(time.RFC3339), flight.FlightDate.Format(time.RFC3339), flight.ArrivalDelay, flight.DepartureDelay)
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
