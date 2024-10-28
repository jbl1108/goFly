package model

import "time"

type FlightData struct {
	IataFlightCode string
	DepartureDelay float64
	ArrivalDelay   float64
	FlightDate     time.Time
}
