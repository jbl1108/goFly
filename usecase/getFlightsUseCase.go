package usecase

import (
	"log"

	"github.com/jbl1108/goFly/usecase/ports"
)

type GetFlightsUseCase struct {
	flightStorage ports.FlightStorage
}

func NewGetFlightsUseCase(flightStorage ports.FlightStorage) *GetFlightsUseCase {
	getFlights := new(GetFlightsUseCase)
	getFlights.flightStorage = flightStorage
	return getFlights
}

func (getFlights *GetFlightsUseCase) Start() error {
	return nil
}

func (getFlights *GetFlightsUseCase) GetFlights() ([]string, error) {
	log.Printf("getFlights")
	return getFlights.flightStorage.GetAllFlights()
}
