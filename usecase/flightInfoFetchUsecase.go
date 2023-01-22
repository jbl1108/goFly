package usecase

type FlightInfoFetchUsecase struct {
	mqtthandler Communicator
}

type FlightData struct {
	IataFlightCode  string
	Departure_delay float64
	Arrival_delay   float64
}

func NewFlightInfoFetcher(mqttHandler Communicator) *FlightInfoFetchUsecase {
	newFlightInfoFetcherUseCase := new(FlightInfoFetchUsecase)
	newFlightInfoFetcherUseCase.mqtthandler = mqttHandler
	return newFlightInfoFetcherUseCase
}

func (newDatahandler FlightInfoFetchUsecase) Start() {

}
