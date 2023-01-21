package flyclient

import (
	"fmt"
	"github.com/jbl1108/goFly/restservice"
	"os"
	"testing"
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

func TestRequest(t *testing.T) {
	result, err := Request("http://localhost:8000/gofly")

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
