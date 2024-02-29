package usecase

import (
	"log"
	"time"

	"github.com/jbl1108/goFly/goFetch/usecase/ports"
)

type InsertFlightUseCase struct {
	flightStorage ports.FlightStorage
}

func NewInsertFlightUseCase(flightStorage ports.FlightStorage) *InsertFlightUseCase {
	insertFlight := new(InsertFlightUseCase)
	insertFlight.flightStorage = flightStorage
	return insertFlight
}

func (insertFlight *InsertFlightUseCase) Start() error {
	return nil
}

func (insertFlight *InsertFlightUseCase) InsertFlight(flight string, flightDate time.Time) error {
	log.Printf("Flight inserted. Flight: %s", flight)
	return insertFlight.flightStorage.StoreFlight(flight)
}
