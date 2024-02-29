package delivery

import (
	"log"
	"time"

	"github.com/jbl1108/goFly/goFetch/usecase"
	"github.com/jbl1108/goFly/goFetch/util"
)

type FligthFetchService struct {
	fligthFetchUseCase usecase.FlightInfoFetchUsecase
	ticker             *time.Ticker
	config             util.Config
}

func NewFligthFetchService(config util.Config, flightFetchUsecase usecase.FlightInfoFetchUsecase) *FligthFetchService {
	ffs := new(FligthFetchService)
	ffs.fligthFetchUseCase = flightFetchUsecase
	ffs.config = config
	return ffs
}

func (ffs *FligthFetchService) Start() error {
	err := ffs.fligthFetchUseCase.Start()
	if err == nil {
		ffs.fligthFetchUseCase.Fetch()
		interval := time.Duration(ffs.config.RestFlightFetchIntervalDays()) * time.Hour * 24
		ffs.ticker = time.NewTicker(interval)
		log.Printf("Starting fetchtimer. Interval %s", interval)
		go func() {
			for range ffs.ticker.C {
				ffs.fligthFetchUseCase.Fetch()
			}
		}()
	}
	return err
}

func (ffs *FligthFetchService) Stop() {
	ffs.ticker.Stop()
}
