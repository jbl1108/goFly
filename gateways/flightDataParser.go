package gateways

import (
	"time"

	"github.com/jbl1108/goFly/model"
)

type FligthDataParser struct {
}

func NewFlightDataParser() *FligthDataParser {
	flightDataParser := new(FligthDataParser)
	return flightDataParser
}

func (m *FligthDataParser) ParseData(message Message) ([]model.FlightData, error) {
	resultMap := message.(map[string]interface{})["flights"].([]interface{})
	var returnValue []model.FlightData
	for _, v := range resultMap {
		valueMap := v.(map[string]interface{})
		var arrival_delay = valueMap["arrival_delay"].(float64)
		var departure_delay = valueMap["departure_delay"].(float64)
		var flight_date_str = valueMap["scheduled_out"].(string)
		var flight_date, _ = time.Parse(time.RFC3339, flight_date_str)
		returnValue = append(returnValue, model.FlightData{
			DepartureDelay: departure_delay, ArrivalDelay: arrival_delay, FlightDate: flight_date})
	}

	return returnValue, nil
}
