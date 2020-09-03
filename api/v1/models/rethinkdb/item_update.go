package mysql

import (
	inp "go-api/api/v1/input"
	"os"

	r "github.com/dancannon/gorethink"
	log "github.com/sirupsen/logrus"
)

func ItemUpdateDB(_p *inp.URLParams, _d inp.ItemData) error {
	log.Info("Calling ItemUpdateDB")
	var err error

	_, err = r.DB(os.Getenv("RDB_ENV") + "_test").Table(os.Getenv("RDB_ENV") + "_test_table").Get((*_p)[inp.QItemID]).Update(map[string]interface{}{
		"stringField": _d.StringField, "listField": _d.ListField, "intField": _d.IntField, "boolField": _d.BoolField}).RunWrite(session)

	if err != nil {
		log.Error("Error updating item data")
		return err
	}

	log.Info("Changes updated correctly, returning")

	return nil
}
