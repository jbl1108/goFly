package gateways

import (
	"time"

	"github.com/jbl1108/goFly/goFetch/model"
	"github.com/jbl1108/goFly/goFetch/util"
)

const DATE_FORMAT = "2006-01-02"

type flightInfoFetcher struct {
	restClient *RestClient
	config     *util.Config
}

func NewFlightInfoFetcher(config *util.Config, restClient *RestClient) *flightInfoFetcher {
	newFetchFlightInfoAdapter := new(flightInfoFetcher)
	newFetchFlightInfoAdapter.config = config
	newFetchFlightInfoAdapter.restClient = restClient
	return newFetchFlightInfoAdapter
}

func (m *flightInfoFetcher) Start() error {
	return nil
}

func (m *flightInfoFetcher) SendFlightRequest(flightCode string, startDate time.Time, endDate time.Time) ([]model.FlightData, error) {
	parser := NewFlightDataParser()

	var request = m.config.FlightInfoRequest() + "/" + flightCode + "?start=" + startDate.Format(DATE_FORMAT) + "&end=" + endDate.Format(DATE_FORMAT)
	response, err := m.restClient.Request(request, map[string]string{"x-apikey": m.config.FlightInfoKey()})

	if err != nil {
		return nil, err
	} else {
		return parser.ParseData(response)
	}
}
