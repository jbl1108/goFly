// main.go
package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jbl1108/goFly/driver"
	"github.com/jbl1108/goFly/restservice"
	"github.com/jbl1108/goFly/usecase"
	"github.com/jbl1108/goFly/util"
)

var newFlightInfoFetcher *usecase.FlightInfoFetchUsecase

func main() {

	/*	const API_KEY string = "938d0e9b-d993-450b-b58a-7ea5798d1066"
		https://iatacodes.org/api/v6/airports?api_key=938d0e9b-d993-450b-b58a-7ea5798d1066&code=BLL
				fmt.Println("goFly")
				result, err := restclient.Request("http://localhost:8000/gofly")
				//result, err := flyclient.Request("https://iatacodes.org/api/v6/airports?api_key=" + API_KEY + "&code=BLL")
				if err != nil {
					fmt.Println("Fetch flight error:" + err.Error())
				} else {
					fmt.Println(result)
				}
	*/

	go restservice.Start()

	var config = util.NewConfig()
	var restClient = driver.NewRestClient()
	var persister = driver.NewRedisDriver(config)
	var mqqtClient = driver.NewMQTTCommunicator(config.MQTTHost() + ":" + strconv.Itoa(config.MQTTPort()))
	var newFetchFlightInfoAdapter = driver.NewFetchFlightInfoAdapter(config, restClient, mqqtClient)
	newFlightInfoFetcher = usecase.NewFlightInfoFetcher(newFetchFlightInfoAdapter, persister)
	persister.StoreString(util.KEY_START_DATE, util.DEFAULT_START_DATE)
	persister.StoreString(util.KEY_END_DATE, util.DEFAULT_END_DATE)
	persister.StoreList(util.KEY_FLIGTH, []string{util.DEFAULT_FLIGHT})
	err := newFlightInfoFetcher.Start()
	if err != nil {
		log.Fatal(err)
	}
	startFetcher()
}

func startFetcher() {

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	mychannel := make(chan bool)

	go func() {
		fmt.Println("Press the Enter Key to stop anytime")
		fmt.Scanln()
		mychannel <- true
	}()
	for {
		select {
		// Case statement
		case <-mychannel:
			return
		case <-ticker.C:
			newFlightInfoFetcher.Fetch()
		}
	}

}
