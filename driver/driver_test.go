package driver

import (
	"fmt"
	"os"
	"testing"

	"github.com/jbl1108/goFly/restservice"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	go restservice.Start()
}

func shutdown() {

}

func TestGoFlyRequest(t *testing.T) {
	var restClient = NewRestClient()

	result, err := restClient.Request("http://localhost:8000/gofly")

	resultMap := result.(map[string]interface{})["response"].([]interface{})[0].(map[string]interface{})

	for k, v := range resultMap {
		fmt.Println(k, "=>", v)
	}

	if err != nil {
		t.Fatal("error:", err.Error())
	}
	if resultMap["code"].(string) != "CDG" {
		t.Error("code error")
	}

}

func TestFlightRequest(t *testing.T) {
	var restClient = NewRestClient()
	result, err1 := restClient.Request("http://localhost:8000/flight")
	fmt.Println(err1)
	if err1 != nil {
		t.Fatal(err1)
	}
	var parser = NewFlightDataParser()
	parsed, err2 := parser.ParseData(result)
	if err2 != nil {
		t.Fail()
	}
	if len(parsed) != 2 {
		t.Fail()
	}
}
