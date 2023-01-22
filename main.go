// main.go
package main

import (
	"github.com/jbl1108/goFly/driver"
	"github.com/jbl1108/goFly/usecase"
)

func main() {

	/*	const API_KEY string = "938d0e9b-d993-450b-b58a-7ea5798d1066"

		fmt.Println("goFly")
		result, err := restclient.Request("http://localhost:8000/gofly")
		//result, err := flyclient.Request("https://iatacodes.org/api/v6/airports?api_key=" + API_KEY + "&code=BLL")
		if err != nil {
			fmt.Println("Fetch flight error:" + err.Error())
		} else {
			fmt.Println(result)
		}
	*/
	var emptyDataParser = usecase.NewEmptyDataParser()
	var mqttClient = driver.NewMQTTCommunicator(emptyDataParser, "flightInfo", "localhost")
	var newFlightInfoFetcher = usecase.NewFlightInfoFetcher(mqttClient)
	newFlightInfoFetcher.Start()
}
