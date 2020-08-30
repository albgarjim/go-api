package mysql

import (
	inp "goggers/api/v1/input"

	out "goggers/api/v1/output"
	"os"

	r "github.com/dancannon/gorethink"

	log "github.com/sirupsen/logrus"
)

func RetrieveItemDB(_p *inp.URLParams) (*out.ItemDataRDB, error) {
	log.Info("RetrieveWishlist call")

	cursor, err := r.DB(os.Getenv("RDB_ENV") + "_test").Table(os.Getenv("RDB_ENV") + "_test_table").Get((*_p)[inp.QItemID]).Run(session)

	if err != nil {
		log.Error("Error retriving wishlist data", err)
		return nil, err
	}

	var item *out.ItemDataRDB
	cursor.One(&item)
	cursor.Close()

	return item, nil
}
