package config

import (
	"github.com/jbl1108/goFly/delivery"
	"github.com/jbl1108/goFly/gateways"
	"github.com/jbl1108/goFly/repositories"
	"github.com/jbl1108/goFly/usecase"
	"github.com/jbl1108/goFly/util"
)

func NewFetchFlightUseCase() *usecase.FlightInfoFetchUsecase {
	var config = util.NewConfig()
	var restClient = gateways.NewRestClient()
	var persister = repositories.NewRedisRepository(config)
	var mqqtClient = delivery.NewMQTTCommunicator(config.MQTTAddr())
	var mqqtFlightPublisher = delivery.NewFlightMQTTPublisher(config, mqqtClient)
	var newFetchFlightInfoAdapter = gateways.NewFlightInfoFetcher(config, restClient)

	//TODO: remove
	persister.StoreString(util.KEY_START_DATE, util.DEFAULT_START_DATE)
	persister.StoreString(util.KEY_END_DATE, util.DEFAULT_END_DATE)
	persister.StoreList(util.KEY_FLIGTH, []string{util.DEFAULT_FLIGHT})

	return usecase.NewFlightInfoFetcher(newFetchFlightInfoAdapter, mqqtFlightPublisher, persister)
}
