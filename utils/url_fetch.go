package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)


func Fetch(_method string, _url string, _header map[string]string, _body map[string]interface{}) (*http.Response, error) {
	// Options: Disable keep-alives, 30sec n/w call timeout.
	log.Info("Making call to ", _url, " site")
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: time.Duration(10 * time.Second),
	}

	bytesRepresentation, err := json.Marshal(_body)
	if err != nil {
		log.Error("Error marshalling body: ", err)
	}

	req, _ := http.NewRequest(_method, _url, bytes.NewBuffer(bytesRepresentation))
	for key, value := range _header {
		req.Header.Add(key, value)
	}

	log.Info(req)

	res, err := client.Do(req)
	if err != nil {
		log.Error("Error performing request: ", err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		log.Info("Call to ", _url, " did not return success", res)
		return nil, errors.New("Network call did not return success")
	}
	return res, nil
}

func FetchRes(_method string, _url string, _header map[string]string, _body map[string]interface{}) (*map[string]interface{}, error) {
	// Options: Disable keep-alives, 30sec n/w call timeout.

	resp, err := Fetch(_method, _url, _header, _body)

	var bodyText []byte
	if bodyText, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Error("Error reading response from call: ", err)
		return nil, err
	}

	var data *map[string]interface{}
	if err = json.Unmarshal([]byte(bodyText), &data); err != nil {
		log.Error("Error processing response: ", err)
		return nil, err
	}

	return data, nil
}


func MakeGetRequest(_url string) map[string]interface{} {
	var err error
	resp, err := http.Get(_url)
	if err != nil {
		log.Error("Error making request to aboutyou", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading body", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		log.Error("Error processing about you response: ", err)
	}

	return data
}