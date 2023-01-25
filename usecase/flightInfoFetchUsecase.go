package usecase

import (
	"log"
	"time"

	"github.com/jbl1108/goFly/util"
	"go.uber.org/multierr"
)

//https://aeroapi.flightaware.com/aeroapi/flights/KL6027?start=2023-01-19T19:59:59Z&end=2023-01-21T19:59:59Z

type FlightInfoFetchUsecase struct {
	communicatior Communicator
	persister     Persistance
}

type FlightData struct {
	IataFlightCode  string
	Departure_delay float64
	Arrival_delay   float64
}

func NewFlightInfoFetcher(communicatior Communicator, persister Persistance) *FlightInfoFetchUsecase {
	newFlightInfoFetcherUseCase := new(FlightInfoFetchUsecase)
	newFlightInfoFetcherUseCase.persister = persister
	newFlightInfoFetcherUseCase.communicatior = communicatior
	return newFlightInfoFetcherUseCase
}

func (fifu *FlightInfoFetchUsecase) Start() {

	ticker := time.NewTicker(3 * time.Second)

	mychannel := make(chan bool)

	// Go function
	go func() {

		select {

		// Case statement
		case <-mychannel:
			return
		case _ = <-ticker.C:
			strStartDate, err1 := fifu.persister.FetchString(util.KEY_START_DATE)
			strEndDate, err2 := fifu.persister.FetchString(util.KEY_END_DATE)

			startDate, err3 := time.Parse(time.RFC3339, strStartDate)
			endDate, err4 := time.Parse(time.RFC3339, strEndDate)
			flights, err5 := fifu.persister.FetchList(util.KEY_FLIGTH)
			errors := multierr.Combine(err1, err2, err3, err4, err5)

			if errors != nil {
				log.Fatalf("error parsing date: %s", errors)
			}

			fifu.fetchFlights(flights, startDate, endDate)
		}
	}()
	time.Sleep(7 * time.Second)

	// Calling Stop() method
	ticker.Stop()

	// Setting the value of channel
	mychannel <- true

}
func (fifu *FlightInfoFetchUsecase) fetchFlights(flightCodes []string, startDate time.Time, endDate time.Time) {

	for _, flightCode := range flightCodes {
		flightData, err := fifu.communicatior.SendFlightRequest(flightCode, startDate, endDate)
		if err == nil {
			fifu.communicatior.PostMessage(flightData)
			log.Printf("Fetched %o flights", len(flightData))
		} else {
			log.Fatalf("Fetch fligth error %s", err.Error())
		}
	}

}
