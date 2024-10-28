package delivery

import (
	"html/template"
	"io"
)

type a struct {
	Flights []string
	Status  string
}

type WebPage struct {
}

func NewWebpage() *WebPage {
	fis := new(WebPage)
	return fis
}

func (wp *WebPage) Generate(wr io.Writer, flights []string, status string) {
	var datas = &a{Flights: flights, Status: status}
	tpl := template.Must(template.ParseFiles("index.html"))
	tpl.Execute(wr, datas)
}
