package config

import (
	"github.com/jbl1108/goFly/delivery"
	"github.com/jbl1108/goFly/gateways"
	"github.com/jbl1108/goFly/repositories"
	"github.com/jbl1108/goFly/usecase"
	"github.com/jbl1108/goFly/usecase/ports"
	"github.com/jbl1108/goFly/util"
)

type applicationFactory struct {
	config                 *util.Config
	keyValueStore          repositories.KeyValueStore
	flightStorage          ports.FlightStorage
	flightInfoFetchUsecase *usecase.FlightInfoFetchUsecase
}

func NewApplication() *applicationFactory {
	app := new(applicationFactory)
	app.config = util.NewConfig()
	app.keyValueStore = repositories.NewRedisRepository(app.config)
	app.flightStorage = repositories.NewFlightRepository(app.keyValueStore)
	app.flightInfoFetchUsecase = app.NewFetchFlightUseCase()
	return app
}

func (app *applicationFactory) NewFlightInputservice() *delivery.FlightInputService {
	return delivery.NewFlightInputService(*app.config, *usecase.NewInsertFlightUseCase(app.flightStorage), *usecase.NewDeleteFlightUseCase(app.flightStorage), *usecase.NewGetFlightsUseCase(app.flightStorage), *app.flightInfoFetchUsecase)
}

func (app *applicationFactory) NewFlightFetchService() *delivery.FligthFetchService {
	return delivery.NewFligthFetchService(*app.config, *app.NewFetchFlightUseCase())
}

func (app *applicationFactory) GetInsertFlightsUseCase() *usecase.GetFlightsUseCase {
	return usecase.NewGetFlightsUseCase(app.flightStorage)
}

func (app *applicationFactory) NewInsertFlightUseCase() *usecase.InsertFlightUseCase {
	return usecase.NewInsertFlightUseCase(app.flightStorage)
}

func (app *applicationFactory) NewDeleteFlightUseCase() *usecase.DeleteFlightUseCase {
	return usecase.NewDeleteFlightUseCase(app.flightStorage)
}

func (app *applicationFactory) NewFetchFlightUseCase() *usecase.FlightInfoFetchUsecase {
	var restClient = gateways.NewRestClient()
	var mqqtClient = gateways.NewMQTTCommunicator(app.config.MQTTAddr())
	var mqqtFlightPublisher = gateways.NewFlightMQTTPublisher(app.config, mqqtClient)
	var newFetchFlightInfoAdapter = gateways.NewFlightInfoFetcher(app.config, restClient)
	return usecase.NewFlightInfoFetcher(newFetchFlightInfoAdapter, mqqtFlightPublisher, app.flightStorage)
}
