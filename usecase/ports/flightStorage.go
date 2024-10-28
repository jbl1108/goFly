package ports

import "time"

type FlightStorage interface {
	StoreStartDate(startDate time.Time) error
	FetchStartDate() (time.Time, error)
	FetchEndDate() (time.Time, error)
	StoreEndDate(startDate time.Time) error
	StoreFlight(flight string) error
	DeleteFlight(flight string) error
	GetAllFlights() ([]string, error)
}
