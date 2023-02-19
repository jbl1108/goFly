// main.go
package main

import (
	"fmt"
	"log"

	"github.com/jbl1108/goFly/config"
	"github.com/jbl1108/goFly/delivery"
	"github.com/jbl1108/goFly/restservice"
	"github.com/jbl1108/goFly/usecase"
	"go.uber.org/multierr"
)

var flightInfoFetcher *usecase.FlightInfoFetchUsecase
var flightInputService *delivery.FlightInputService

func main() {

	go restservice.Start()
	app := config.NewApplication()
	flightInfoFetcher = app.NewFetchFlightUseCase()
	flightInputService = app.NewFlightInputservice()
	err1 := flightInfoFetcher.Start()
	err2 := flightInputService.Start()
	err := multierr.Append(err1, err2)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Press the Enter Key to stop anytime")
		fmt.Scanln()
		flightInfoFetcher.Stop()
	}

}
