package main

import (
	"io"
	"net/http"
)

func Start() {
	http.HandleFunc("/gofly", gofly)
	http.HandleFunc("/going", going)
	http.ListenAndServe(":8000", nil)
}

func gofly(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, getJson())
}

func going(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "going!")
}

func main() {
	Start()
}

func getJson() (b string) {
	b = `{"response":[{"code":"CDG","city_code":"PAR","country_code":"FR","name":"Charles De Gaulle","alternatenames":"Paris-Charles-de-Gaulle,Charles De Gaulle,Ch. De Gaulle,Ch. De Gaulle,Ch. De Gaulle,Charles Degaulle","lat":49.003197,"lng":2.567023,"timezone":"Europe/Paris","gmt":120,"popularity":182,"is_rail_road":0,"is_bus_station":0,"icao":"LFPG","phone":"(01) 4862 121","site":"http://www.cdgfacile.com","geoname_id":"6269554","routes":42}]}`
	return
}
