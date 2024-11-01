package restservice

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/gofly", gofly)
	r.HandleFunc("/flights/{id}", flight)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":8000", r)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Print("not found: " + r.RequestURI)
}

func gofly(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, getGoFlyJson())
}

func flight(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, getFlightJson())
}

func getGoFlyJson() (b string) {
	b = `{"response":[{"code":"CDG","city_code":"PAR","country_code":"FR","name":"Charles De Gaulle","alternatenames":"Paris-Charles-de-Gaulle,Charles De Gaulle,Ch. De Gaulle,Ch. De Gaulle,Ch. De Gaulle,Charles Degaulle","lat":49.003197,"lng":2.567023,"timezone":"Europe/Paris","gmt":120,"popularity":182,"is_rail_road":0,"is_bus_station":0,"icao":"LFPG","phone":"(01) 4862 121","site":"http://www.cdgfacile.com","geoname_id":"6269554","routes":42}]}`
	return
}
func getFlightJson() (b string) {
	b = `{
		"flights": [
			{
				"ident": "DAL57",
				"ident_icao": "DAL57",
				"ident_iata": "DL57",
				"fa_flight_id": "DAL57-1674104977-fa-0001",
				"operator": "DAL",
				"operator_icao": "DAL",
				"operator_iata": "DL",
				"flight_number": "57",
				"registration": "N812NW",
				"atc_ident": null,
				"inbound_fa_flight_id": "DAL132-1674018614-fa-0001",
				"codeshares": [
					"AFR5594",
					"KLM6027",
					"VIR3942"
				],
				"codeshares_iata": [
					"AF5594",
					"KL6027",
					"VS3942"
				],
				"blocked": false,
				"diverted": false,
				"cancelled": false,
				"position_only": false,
				"origin": {
					"code": "EHAM",
					"code_icao": "EHAM",
					"code_iata": "AMS",
					"code_lid": null,
					"timezone": "Europe/Amsterdam",
					"name": "Amsterdam Schiphol",
					"city": "Amsterdam",
					"airport_info_url": "/airports/EHAM"
				},
				"destination": {
					"code": "KSLC",
					"code_icao": "KSLC",
					"code_iata": "SLC",
					"code_lid": "SLC",
					"timezone": "America/Denver",
					"name": "Salt Lake City Intl",
					"city": "Salt Lake City",
					"airport_info_url": "/airports/KSLC"
				},
				"departure_delay": 60,
				"arrival_delay": -1260,
				"filed_ete": 35760,
				"foresight_predictions_available": true,
				"scheduled_out": "2023-01-21T10:15:00Z",
				"estimated_out": "2023-01-21T10:15:00Z",
				"actual_out": "2023-01-21T10:16:00Z",
				"scheduled_off": "2023-01-21T10:25:00Z",
				"estimated_off": "2023-01-21T10:59:19Z",
				"actual_off": "2023-01-21T10:59:19Z",
				"scheduled_on": "2023-01-21T20:21:00Z",
				"estimated_on": "2023-01-21T20:29:32Z",
				"actual_on": "2023-01-21T20:29:32Z",
				"scheduled_in": "2023-01-21T20:55:00Z",
				"estimated_in": "2023-01-21T20:37:00Z",
				"actual_in": "2023-01-21T20:34:00Z",
				"progress_percent": 100,
				"status": "Arrived / Gate Arrival",
				"aircraft_type": "A333",
				"route_distance": 4982,
				"filed_airspeed": 463,
				"filed_altitude": 340,
				"route": "BERGI L602 AMGOD L602 SUPUR P1 GIGUL N44 UPNAL RIXUN 660000N/0100000W 710000N/0200000W 740000N/0400000W 740000N/0600000W MEDPA 720000N/0700000W 690000N/0800000W 650000N/0900000W 600000N/0970000W 560000N/1000000W 550000N/1010000W 500000N/1050000W GGW BIL DNW JAC NORDK6",
				"baggage_claim": null,
				"seats_cabin_business": 34,
				"seats_cabin_coach": 199,
				"seats_cabin_first": 17,
				"gate_origin": "E2",
				"gate_destination": "A21",
				"terminal_origin": null,
				"terminal_destination": null,
				"type": "Airline"
			},
			{
				"ident": "DAL57",
				"ident_icao": "DAL57",
				"ident_iata": "DL57",
				"fa_flight_id": "DAL57-1674018613-fa-0000",
				"operator": "DAL",
				"operator_icao": "DAL",
				"operator_iata": "DL",
				"flight_number": "57",
				"registration": "N822NW",
				"atc_ident": null,
				"inbound_fa_flight_id": "DAL134-1673932428-fa-0000",
				"codeshares": [
					"AFR5594",
					"KLM6027",
					"VIR3942"
				],
				"codeshares_iata": [
					"AF5594",
					"KL6027",
					"VS3942"
				],
				"blocked": false,
				"diverted": false,
				"cancelled": false,
				"position_only": false,
				"origin": {
					"code": "EHAM",
					"code_icao": "EHAM",
					"code_iata": "AMS",
					"code_lid": null,
					"timezone": "Europe/Amsterdam",
					"name": "Amsterdam Schiphol",
					"city": "Amsterdam",
					"airport_info_url": "/airports/EHAM"
				},
				"destination": {
					"code": "KSLC",
					"code_icao": "KSLC",
					"code_iata": "SLC",
					"code_lid": "SLC",
					"timezone": "America/Denver",
					"name": "Salt Lake City Intl",
					"city": "Salt Lake City",
					"airport_info_url": "/airports/KSLC"
				},
				"departure_delay": 0,
				"arrival_delay": 1200,
				"filed_ete": 36900,
				"foresight_predictions_available": true,
				"scheduled_out": "2023-01-20T10:20:00Z",
				"estimated_out": "2023-01-20T10:20:00Z",
				"actual_out": "2023-01-20T10:20:00Z",
				"scheduled_off": "2023-01-20T10:30:00Z",
				"estimated_off": "2023-01-20T10:53:22Z",
				"actual_off": "2023-01-20T10:53:22Z",
				"scheduled_on": "2023-01-20T20:45:00Z",
				"estimated_on": "2023-01-20T21:09:41Z",
				"actual_on": "2023-01-20T21:09:41Z",
				"scheduled_in": "2023-01-20T20:55:00Z",
				"estimated_in": "2023-01-20T21:15:00Z",
				"actual_in": "2023-01-20T21:15:00Z",
				"progress_percent": 100,
				"status": "Arrived / Gate Arrival",
				"aircraft_type": "A333",
				"route_distance": 4982,
				"filed_airspeed": 471,
				"filed_altitude": 340,
				"route": "BERGI L602 AMGOD L602 NALAX L46 ODNEK N110 ERKIT L602 TLA MIMKU BEGID BILTO 570000N/0200000W 590000N/0300000W 600000N/0400000W 600000N/0500000W PIDSO 590000N/0700000W 570000N/0800000W 550000N/0870000W 540000N/0900000W 500000N/0980000W MOT SWTHN JAC NORDK6",
				"baggage_claim": null,
				"seats_cabin_business": 30,
				"seats_cabin_coach": 153,
				"seats_cabin_first": 18,
				"gate_origin": "D49",
				"gate_destination": "A25",
				"terminal_origin": null,
				"terminal_destination": null,
				"type": "Airline"
			}
		],
		"links": null,
		"num_pages": 1
	}`
	return
}
