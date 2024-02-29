package repositories

import (
	"time"

	"github.com/jbl1108/goFly/goFetch/util"
)

type FlightRepository struct {
	keyValueStore KeyValueStore
}

func NewFlightRepository(keyValueStore KeyValueStore) *FlightRepository {
	frs := new(FlightRepository)
	frs.keyValueStore = keyValueStore
	return frs
}

func (frs *FlightRepository) StoreStartDate(startDate time.Time) error {
	strStartDate := startDate.Format(time.RFC3339)
	return frs.keyValueStore.StoreString(util.KEY_START_DATE, strStartDate)
}
func (frs *FlightRepository) FetchStartDate() (time.Time, error) {
	startDate, err := frs.keyValueStore.FetchString(util.KEY_START_DATE)
	if err == nil {
		return time.Parse(time.RFC3339, startDate)
	} else {
		return time.Now(), err
	}

}
func (frs *FlightRepository) StoreEndDate(endDate time.Time) error {
	strEndDate := endDate.Format(time.RFC3339)
	return frs.keyValueStore.StoreString(util.KEY_END_DATE, strEndDate)
}
func (frs *FlightRepository) FetchEndDate() (time.Time, error) {
	endDate, err := frs.keyValueStore.FetchString(util.KEY_END_DATE)
	if err == nil {
		return time.Parse(time.RFC3339, endDate)
	} else {
		return time.Now(), err
	}
}

func (frs *FlightRepository) StoreFlight(flight string) error {
	var aFlight []string
	aFlight = append(aFlight, flight)
	return frs.keyValueStore.AppendToList(util.KEY_FLIGTH, aFlight)
}

func (frs *FlightRepository) DeleteFlight(flight string) error {
	allFlights, err := frs.GetAllFlights()
	var newFlightList []string
	if err == nil {
		for _, value := range allFlights {
			if value != flight {
				newFlightList = append(newFlightList, value)
			}
		}
		return frs.keyValueStore.StoreList(util.KEY_FLIGTH, newFlightList)
	}
	return err

}
func (frs *FlightRepository) GetAllFlights() ([]string, error) {
	return frs.keyValueStore.FetchList(util.KEY_FLIGTH)
}
