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
	config        *util.Config
	keyValueStore repositories.KeyValueStore
	flightStorage ports.FlightStorage
}

func NewApplication() *applicationFactory {
	app := new(applicationFactory)
	app.config = util.NewConfig()
	app.keyValueStore = repositories.NewRedisRepository(app.config)
	app.flightStorage = repositories.NewFlightRepository(app.keyValueStore)
	return app
}

func (app *applicationFactory) NewFlightInputservice() *delivery.FlightInputService {
	return delivery.NewFlightInputService(*app.config, *usecase.NewInsertFlightUseCase(app.flightStorage), *usecase.NewDeleteFlightUseCase(app.flightStorage), *usecase.NewGetFlightsUseCase(app.flightStorage))
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
	var mqqtClient = delivery.NewMQTTCommunicator(app.config.MQTTAddr())
	var mqqtFlightPublisher = delivery.NewFlightMQTTPublisher(app.config, mqqtClient)
	var newFetchFlightInfoAdapter = gateways.NewFlightInfoFetcher(app.config, restClient)

	//TODO: remove
	app.keyValueStore.StoreString(util.KEY_START_DATE, util.DEFAULT_START_DATE)
	app.keyValueStore.StoreString(util.KEY_END_DATE, util.DEFAULT_END_DATE)
	app.keyValueStore.StoreList(util.KEY_FLIGTH, []string{util.DEFAULT_FLIGHT})
	return usecase.NewFlightInfoFetcher(newFetchFlightInfoAdapter, mqqtFlightPublisher, app.flightStorage)
}
