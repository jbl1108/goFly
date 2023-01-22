package driver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message interface {
}

type restClient struct {
}

func NewRestClient() *restClient {
	restClient := new(restClient)
	return restClient
}

func (m *restClient) Request(url string) (result Message, err error) {
	var resp *http.Response
	var body []byte

	resp, err = http.Get(url)

	if err != nil {
		defer resp.Body.Close()
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	} else if resp.StatusCode > 299 || resp.StatusCode < 200 {
		err = fmt.Errorf("error: response returned: %s", resp.Status)
	} else {
		result = make(map[string]string, 0)
		err = json.Unmarshal(body, &result)
	}
	return
}
