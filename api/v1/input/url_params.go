package input

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

const (
	QItemID   = "itemId"
	QItemList = "itemList"
)

var list = []string{QItemID, QItemList}

type URLParams = map[string]interface{}

func ExtractParameters(_r *http.Request) (*URLParams, error) {
	vars := mux.Vars(_r)
	params := &URLParams{}

	for _, key := range list {
		val := valid[key](_r.URL.Query().Get(key))
		if val != "" {
			(*params)[key] = val
		}
	}

	for key, val := range vars {
		(*params)[key] = valid[key](val)
	}

	for key, val := range *params {
		log.Info("query: ", key, val)
	}

	return params, nil
}
