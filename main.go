// main.go
package main

import (
	"fmt"
	"log"

	"github.com/jbl1108/goFly/config"
	"github.com/jbl1108/goFly/restservice"
	"github.com/jbl1108/goFly/usecase"
)

var newFlightInfoFetcher *usecase.FlightInfoFetchUsecase

func main() {

	go restservice.Start()

	newFlightInfoFetcher = config.NewFetchFlightUseCase()
	err := newFlightInfoFetcher.Start()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Press the Enter Key to stop anytime")
		fmt.Scanln()
		newFlightInfoFetcher.Stop()
	}

}
