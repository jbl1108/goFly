package ports

import "github.com/jbl1108/goFly/goFetch/model"

type FlightPublisher interface {
	Start() error
	PostMessage(message []model.FlightData) error
}
