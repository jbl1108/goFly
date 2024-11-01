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
	log             *log.Logger
}

func NewFlightInfoFetcher(flightFetcher ports.FlightFetcher, flightPublisher ports.FlightPublisher, flightStorage ports.FlightStorage) *FlightInfoFetchUsecase {
	newFlightInfoFetcherUseCase := new(FlightInfoFetchUsecase)
	newFlightInfoFetcherUseCase.flightStorage = flightStorage
	newFlightInfoFetcherUseCase.flightFetcher = flightFetcher
	newFlightInfoFetcherUseCase.flightPublisher = flightPublisher
	newFlightInfoFetcherUseCase.log = log.Default()
	newFlightInfoFetcherUseCase.log.SetPrefix("FlightInfoFetchUsecase : ")
	return newFlightInfoFetcherUseCase
}

func (fifu *FlightInfoFetchUsecase) Start() error {
	_, err1 := fifu.flightStorage.FetchStartDate()
	if err1 != nil {
		fifu.flightStorage.StoreStartDate(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	}
	_, err2 := fifu.flightStorage.FetchEndDate()

	if err2 != nil {
		fifu.flightStorage.StoreEndDate(time.Now())
	}

	return multierr.Combine(fifu.flightPublisher.Start(), fifu.flightFetcher.Start())
}

func (fifu *FlightInfoFetchUsecase) Fetch() error {
	startDate, err1 := fifu.flightStorage.FetchStartDate()
	endDate, err2 := fifu.flightStorage.FetchEndDate()

	flights, err3 := fifu.flightStorage.GetAllFlights()
	errors := multierr.Combine(err1, err2, err3)

	if errors != nil {
		log.Fatalf("error parsing date: %s", errors)
	}
	fifu.fetchFlights(flights, startDate, endDate)

	fifu.flightStorage.StoreStartDate(time.Now().Add(time.Duration(-24) * time.Hour))
	fifu.flightStorage.StoreEndDate(time.Now())
	return errors
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
					fifu.log.Printf("Error posting message: %s", err.Error())
				}

			}
			fifu.log.Printf("Fetched %o flights", len(flightData))
		} else {
			fifu.log.Printf("Fetch fligth error %s", err.Error())
		}

	}

}
