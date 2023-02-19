package delivery

import (
	"html/template"
	"io"
)

type a struct {
	Flights []string
}

type WebPage struct {
}

func NewWebpage() *WebPage {
	fis := new(WebPage)
	return fis
}

func (wp *WebPage) Generate(wr io.Writer, data any) {

	var datas = &a{Flights: []string{"One", "Two", "Three"}}

	tpl := template.Must(template.ParseFiles("delivery/index.html"))
	tpl.Execute(wr, datas)
}
