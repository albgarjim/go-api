package mysql

import (
	"os"

	inp "goggers/api/v1/input"

	r "github.com/dancannon/gorethink"
	log "github.com/sirupsen/logrus"
)

func ItemInsertDB(_d inp.ItemData) error {
	var err error
	log.Info("Calling CreateAikoneFacebookProfile")
	log.Info("Creating facebook profile")

	_, err = r.DB(os.Getenv("RDB_ENV") + "_test").Table(os.Getenv("RDB_ENV") + "_test_table").Insert(_d).RunWrite(session)

	if err != nil {
		log.Error("Error inserting facebook profile")
		return err
	}

	log.Info("Aikone profile created, returning")

	return nil
}
