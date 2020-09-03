package mysql

import (
	"os"

	inp "go-api/api/v1/input"

	r "github.com/dancannon/gorethink"
	log "github.com/sirupsen/logrus"
)

func ItemInsertDB(_d inp.ItemData) error {
	var err error
	log.Info("Creating product profile")

	_, err = r.DB(os.Getenv("RDB_ENV") + "_test").Table(os.Getenv("RDB_ENV") + "_test_table").Insert(_d).RunWrite(session)

	if err != nil {
		log.Error("Error inserting item in db")
		return err
	}

	log.Info("Item created, returning")

	return nil
}
