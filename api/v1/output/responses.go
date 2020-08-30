package output

import (
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"

	log "github.com/sirupsen/logrus"
)

//RespondWithError returns an error specified by the _code parameter which corresponds to the http error
func RespondWithError(_w http.ResponseWriter, _err error) {
	code, message := ErrorCodeMessage(_err)
	log.Info("Responding with error: ", message)

	resp := make(map[string]interface{})
	resp["timestamp"] = time.Now()
	resp["code"] = code
	resp["error"] = message
	resp["Content-Type"] = "application/json"

	var json = jsoniter.ConfigFastest

	response, _ := json.Marshal(&resp)

	log.Info("Returning json response")

	_w.Header().Set("Content-Type", "application/json")
	_w.WriteHeader(code)
	_w.Write(response)
}

/*
RespondWithJSON writes the json response to return to the user, it contains the following fields

- timestamp: time at the request was made

- data: data of the request

- Content-Type: format of the response
*/
func RespondWithJSON(_w http.ResponseWriter, _code int, payload interface{}) {
	resp := make(map[string]interface{})
	resp["timestamp"] = time.Now()
	resp["data"] = payload
	resp["Content-Type"] = "application/json"

	var json = jsoniter.ConfigFastest

	response, _ := json.Marshal(&resp)

	log.Info("Returning json response")

	_w.Header().Set("Content-Type", "application/json")
	_w.WriteHeader(_code)
	_w.Write(response)
}
