package gateways

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message interface {
}

type RestClient struct {
}

func NewRestClient() *RestClient {
	restClient := new(RestClient)
	return restClient
}

func (m *RestClient) Request(url string, header map[string]string) (result Message, err error) {
	var resp *http.Response
	var body []byte

	client := http.Client{}
	log.Printf("request to: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err = client.Do(req)

	if err != nil {
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
