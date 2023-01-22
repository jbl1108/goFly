package driver

import (
	"fmt"

	"github.com/jbl1108/goFly/usecase"
)

type FligthDataParser struct {
}

func NewFlightDataParser() *FligthDataParser {
	flightDataParser := new(FligthDataParser)
	return flightDataParser
}

func (m *FligthDataParser) ParseData(message Message) ([]usecase.FlightData, error) {
	resultMap := message.(map[string]interface{})["flights"].([]interface{})
	fmt.Println(resultMap)
	var returnValue []usecase.FlightData
	for _, v := range resultMap {
		valueMap := v.(map[string]interface{})
		var arrival_delay = valueMap["arrival_delay"].(float64)
		var departure_delay = valueMap["departure_delay"].(float64)
		returnValue = append(returnValue, usecase.FlightData{
			Departure_delay: departure_delay, Arrival_delay: arrival_delay})

	}

	return returnValue, nil
}
