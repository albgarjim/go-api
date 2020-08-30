package input

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

var valid = map[string]func(string) interface{}{
	QItemID:   valItemID,
	QItemList: valItemList,
}

func valItemID(_val string) interface{} {
	if _val == "" {
		log.Info("Filtered item id")
		return ""
	}
	return _val
}

func valItemList(_val string) interface{} {
	log.Info("idea list: ", _val)
	if _val == "" {
		return [...]string{""}
	}
	return strings.Split(_val, " ")
}
