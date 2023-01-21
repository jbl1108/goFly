package flyclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message interface {
}

func Request(url string) (result Message, err error) {
	var resp *http.Response
	var body []byte
	result = make(map[string]string, 0)
	resp, err = http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		err = json.Unmarshal(body, &result)
	}
	return
}
