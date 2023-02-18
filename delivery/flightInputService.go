package delivery

import (
	"io"
	"log"
	"net/http"
	"time"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/jbl1108/goFly/usecase"
	"github.com/jbl1108/goFly/util"
)

type FlightInputService struct {
	config              util.Config
	insertFlightUsecase usecase.InsertFlightUseCase
	deleteFlightUsecase usecase.DeleteFlightUseCase
	getFlightsUseCase   usecase.GetFlightsUseCase
}

func NewFlightInputService(config util.Config, insertFlightUsecase usecase.InsertFlightUseCase, deleteFlightUsecase usecase.DeleteFlightUseCase, getFlightsUseCase usecase.GetFlightsUseCase) *FlightInputService {
	fis := new(FlightInputService)
	fis.config = config
	fis.insertFlightUsecase = insertFlightUsecase
	fis.deleteFlightUsecase = deleteFlightUsecase
	fis.getFlightsUseCase = getFlightsUseCase
	return fis
}

func (fis *FlightInputService) Start() error {
	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/flights", fis.newFlight).Methods("POST")
		r.HandleFunc("/flights/{id}", fis.delFlight).Methods("DELETE")
		r.HandleFunc("/flights", fis.getFlights).Methods("GET")
		r.NotFoundHandler = http.HandlerFunc(fis.notFound)
		err := http.ListenAndServe(fis.config.RestServiceAddress(), r)
		if err != nil {
			log.Fatal(err)
		}

	}()
	return nil
}

func (fis *FlightInputService) notFound(w http.ResponseWriter, r *http.Request) {
	log.Print("not found: " + r.RequestURI)
}

func (fis *FlightInputService) newFlight(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseForm()
	flight := r.Form.Get("flight")
	if flight == "" {
		io.WriteString(w, "flight parameter missing!")
	} else {
		if err1 == nil {
			err2 := fis.insertFlightUsecase.InsertFlight(flight, time.Now())
			if err2 == nil {
				io.WriteString(w, "OK")
			} else {
				io.WriteString(w, err2.Error())
			}
		} else {
			io.WriteString(w, err1.Error())
		}
	}
}

func (fis *FlightInputService) delFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flight := vars["id"]
	if flight == "" {
		io.WriteString(w, "flight parameter missing!")
	} else {
		err := fis.deleteFlightUsecase.DeleteFlight(flight)
		if err == nil {
			io.WriteString(w, "OK")
		} else {
			io.WriteString(w, err.Error())
		}

	}
}

func (fis *FlightInputService) getFlights(w http.ResponseWriter, r *http.Request) {
	log.Print("getFlights")
	flights, err := fis.getFlightsUseCase.GetFlights()
	if err == nil {
		io.WriteString(w, fis.toJson(flights))
	} else {
		io.WriteString(w, err.Error())
	}
}

func (fis *FlightInputService) toJson(flights []string) string {
	encjson, _ := json.Marshal(flights)
	return string(encjson)
}
