package mysql

import (
	inp "go-api/api/v1/input"
	"os"

	r "github.com/dancannon/gorethink"
	log "github.com/sirupsen/logrus"
)

func ItemDeleteDB(_p *inp.URLParams) error {
	log.Info("Calling ItemDeleteDB")

	if _, err := r.DB(os.Getenv("RDB_ENV") + "_test").Table(os.Getenv("RDB_ENV") + "_test_table").Get((*_p)[inp.QItemID]).Delete().RunWrite(session); err != nil {
		log.Error("Error ideleting item from db", err)
		return err
	}

	log.Info("Data removed correctly, returning")

	return nil
}
