package usecase

type Communicator interface {
	RegisterMessageListener(topic string, listener func(message string))
	PostMessage(message string, receiver string) (error, string)
	SendFlightRequest(request string, response FlightData) ([]FlightData, error)
}
