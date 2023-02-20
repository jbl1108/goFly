// main.go
package main

import (
	"fmt"
	"log"

	"github.com/jbl1108/goFly/config"
	"github.com/jbl1108/goFly/delivery"
	"go.uber.org/multierr"
)

var flightInfoFetchService *delivery.FligthFetchService
var flightInputService *delivery.FlightInputService

func main() {

	//go restservice.Start()
	app := config.NewApplication()
	flightInfoFetchService = app.NewFlightFetchService()
	flightInputService = app.NewFlightInputservice()
	err1 := flightInfoFetchService.Start()
	err2 := flightInputService.Start()
	err := multierr.Append(err1, err2)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Press the Enter Key to stop anytime")
		fmt.Scanln()
		flightInfoFetchService.Stop()
	}

}
