package usecase

import (
	"log"
	"time"

	"github.com/jbl1108/goFly/usecase/ports"
	"go.uber.org/multierr"
)

type FlightInfoFetchUsecase struct {
	flightPublisher ports.FlightPublisher
	flightFetcher   ports.FlightFetcher
	flightStorage   ports.FlightStorage
	ticker          *time.Ticker
}

func NewFlightInfoFetcher(flightFetcher ports.FlightFetcher, flightPublisher ports.FlightPublisher, flightStorage ports.FlightStorage) *FlightInfoFetchUsecase {
	newFlightInfoFetcherUseCase := new(FlightInfoFetchUsecase)
	newFlightInfoFetcherUseCase.flightStorage = flightStorage
	newFlightInfoFetcherUseCase.flightFetcher = flightFetcher
	newFlightInfoFetcherUseCase.flightPublisher = flightPublisher
	return newFlightInfoFetcherUseCase
}

func (fifu *FlightInfoFetchUsecase) Start() error {
	var err = multierr.Combine(fifu.flightPublisher.Start(), fifu.flightFetcher.Start())
	if err == nil {
		fifu.ticker = time.NewTicker(20 * time.Second)
		go func() {
			for range fifu.ticker.C {
				fifu.Fetch()
			}
		}()
	}
	return err
}

func (fifu *FlightInfoFetchUsecase) Stop() {
	fifu.ticker.Stop()
}

func (fifu *FlightInfoFetchUsecase) Fetch() {
	startDate, err1 := fifu.flightStorage.FetchStartDate()
	endDate, err2 := fifu.flightStorage.FetchEndDate()
	flights, err3 := fifu.flightStorage.GetAllFlights()
	errors := multierr.Combine(err1, err2, err3)

	if errors != nil {
		log.Fatalf("error parsing date: %s", errors)
	}
	fifu.fetchFlights(flights, startDate, endDate)
}

func (fifu *FlightInfoFetchUsecase) fetchFlights(flightCodes []string, startDate time.Time, endDate time.Time) {

	for _, flightCode := range flightCodes {
		flightData, err := fifu.flightFetcher.SendFlightRequest(flightCode, startDate, endDate)
		for k := range flightData {
			flightData[k].IataFlightCode = flightCode
		}
		if err == nil {
			if len(flightData) != 0 {
				err = fifu.flightPublisher.PostMessage(flightData)
				if err != nil {
					log.Fatalf("Error posting message: %s", err.Error())
				}

			}
			log.Printf("Fetched %o flights", len(flightData))
		} else {
			log.Fatalf("Fetch fligth error %s", err.Error())
		}

	}

}
