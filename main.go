// main.go
package main

import (
	"encoding/json"

	"fmt"

	"github.com/incoding/gofly/flyclient"
	"github.com/incoding/util"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	var m Message
	fmt.Println(stringutil.Reverse("Hello World!"))
	result := flyclient.Request("http://localhost:8000/gofly")
	err := json.Unmarshal(result, &m)
	fmt.Println(err.Error())
}
