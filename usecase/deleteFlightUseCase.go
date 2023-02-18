package usecase

import (
	"log"

	"github.com/jbl1108/goFly/usecase/ports"
)

type DeleteFlightUseCase struct {
	flightStorage ports.FlightStorage
}

func NewDeleteFlightUseCase(flightStorage ports.FlightStorage) *DeleteFlightUseCase {
	deleteFlight := new(DeleteFlightUseCase)
	deleteFlight.flightStorage = flightStorage
	return deleteFlight
}

func (deleteFlight *DeleteFlightUseCase) Start() error {
	return nil
}

func (deleteFlight *DeleteFlightUseCase) DeleteFlight(flight string) error {
	log.Printf("Flight deleted. Flight: %s", flight)
	return deleteFlight.flightStorage.DeleteFlight(flight)
}
