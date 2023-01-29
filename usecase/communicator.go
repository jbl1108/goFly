package usecase

import "time"

type Communicator interface {
	Start() error
	PostMessage(message []FlightData) error
	SendFlightRequest(flightCode string, startDate time.Time, endDate time.Time) ([]FlightData, error)
}
