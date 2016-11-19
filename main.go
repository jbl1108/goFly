// main.go
package main

import (
	"fmt"

	"github.com/incoding/gofly/flyclient"
)

func main() {

	const API_KEY string = "938d0e9b-d993-450b-b58a-7ea5798d1066"

	fmt.Println("goFly")
	//result, err := flyclient.Request("http://localhost:8000/gofly")
	result, err := flyclient.Request("https://iatacodes.org/api/v6/airports?api_key=" + API_KEY + "&code=BLL")
	if err != nil {
		fmt.Println("Fetch flight error:" + err.Error())
	} else {
		fmt.Println(result)
	}

}
