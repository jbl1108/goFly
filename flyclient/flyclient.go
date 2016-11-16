package flyclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Request(url string) []byte {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}
