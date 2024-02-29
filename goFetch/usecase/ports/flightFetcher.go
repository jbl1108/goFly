package ports

import (
	"time"

	"github.com/jbl1108/goFly/goFetch/model"
)

type FlightFetcher interface {
	Start() error
	SendFlightRequest(flightCode string, startDate time.Time, endDate time.Time) ([]model.FlightData, error)
}
