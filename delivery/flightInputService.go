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
	config                 util.Config
	insertFlightUsecase    usecase.InsertFlightUseCase
	deleteFlightUsecase    usecase.DeleteFlightUseCase
	getFlightsUseCase      usecase.GetFlightsUseCase
	fetchFlightInfoUseCase usecase.FlightInfoFetchUsecase
	webPage                *WebPage
}

func NewFlightInputService(config util.Config, insertFlightUsecase usecase.InsertFlightUseCase, deleteFlightUsecase usecase.DeleteFlightUseCase, getFlightsUseCase usecase.GetFlightsUseCase, fetchFlightInfoUseCase usecase.FlightInfoFetchUsecase) *FlightInputService {
	fis := new(FlightInputService)
	fis.config = config
	fis.insertFlightUsecase = insertFlightUsecase
	fis.deleteFlightUsecase = deleteFlightUsecase
	fis.getFlightsUseCase = getFlightsUseCase
	fis.fetchFlightInfoUseCase = fetchFlightInfoUseCase
	fis.webPage = NewWebpage()
	return fis
}

func (fis *FlightInputService) Start() error {
	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/flights", fis.newFlight).Methods("POST")
		r.HandleFunc("/delflight", fis.delFlight).Methods("POST")
		r.HandleFunc("/fetch", fis.fetchFlight).Methods("POST")
		r.HandleFunc("/flights", fis.getFlights).Methods("GET")
		r.HandleFunc("/", fis.showWebPage).Methods("GET")
		r.HandleFunc("", fis.showWebPage).Methods("GET")
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
	io.WriteString(w, "Page Not found: "+r.RequestURI)
}

func (fis *FlightInputService) fetchFlight(w http.ResponseWriter, r *http.Request) {
	err := fis.fetchFlightInfoUseCase.Fetch()
	var status = ""
	if(err != nil){
		status = err.Error()
	}
	fis.generateWebPage(w, status)
}

func (fis *FlightInputService) newFlight(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseForm()
	flight := r.Form.Get("flight")
	if flight == "" {
		fis.generateWebPage(w, "flight parameter missing!")
	} else {
		if err1 == nil {
			err2 := fis.insertFlightUsecase.InsertFlight(flight, time.Now())
			if err2 == nil {
				fis.generateWebPage(w, "flightAdded")
			} else {
				fis.generateWebPage(w, err2.Error())
			}
		} else {
			fis.generateWebPage(w, err1.Error())
		}
	}
}

func (fis *FlightInputService) delFlight(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseForm()
	flight := r.Form.Get("flight")
	if flight == "" {
		fis.generateWebPage(w, "flight parameter missing!")
	} else {
		if err1 == nil {
			err2 := fis.deleteFlightUsecase.DeleteFlight(flight)
			if err2 == nil {
				fis.generateWebPage(w, "flightDeleted")
			} else {
				fis.generateWebPage(w, err2.Error())
			}
		} else {
			fis.generateWebPage(w, err1.Error())
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

func (fis *FlightInputService) showWebPage(w http.ResponseWriter, r *http.Request) {
	log.Print("showWebPage")
	fis.generateWebPage(w, "")
}

func (fis *FlightInputService) generateWebPage(w io.Writer, status string) {
	flights, err := fis.getFlightsUseCase.GetFlights()
	if err == nil {
		fis.webPage.Generate(w, flights, status)
	} else {
		io.WriteString(w, err.Error())
	}
}

func (fis *FlightInputService) toJson(flights []string) string {
	encjson, _ := json.Marshal(flights)
	return string(encjson)
}
