package gateways

import (
	"fmt"
	"log"
	"time"

	"github.com/jbl1108/goFly/goFetch/model"
	"github.com/jbl1108/goFly/goFetch/util"
	"go.uber.org/multierr"
)

type flightMQTTPublisher struct {
	mqttCommunicator *MQTTCommunicator
	config           *util.Config
}

func NewFlightMQTTPublisher(config *util.Config, mqttCommunicator *MQTTCommunicator) *flightMQTTPublisher {
	flightPublisher := new(flightMQTTPublisher)
	flightPublisher.config = config
	flightPublisher.mqttCommunicator = mqttCommunicator
	return flightPublisher
}

func (m *flightMQTTPublisher) Start() error {
	return m.mqttCommunicator.Start()
}

func (m *flightMQTTPublisher) PostMessage(message []model.FlightData) error {
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
func (m *flightMQTTPublisher) generateJson(flight model.FlightData) string {
	return fmt.Sprintf("{\"time\" : \"%s\" , \"text_flightDate\" : \"%s\" , \"arrivalDelay\" : %f , \"departureDelay\" : %f }", flight.FlightDate.Format(time.RFC3339), flight.FlightDate.Format(time.RFC3339), flight.ArrivalDelay, flight.DepartureDelay)
}
