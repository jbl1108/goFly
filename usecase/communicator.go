package usecase

import "time"

type Communicator interface {
	PostMessage(message []FlightData) error
	SendFlightRequest(flightCode string, startDate time.Time, endDate time.Time) ([]FlightData, error)
}
