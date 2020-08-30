package mysql

import (
	"fmt"
	"os"
	"time"

	r "github.com/dancannon/gorethink"
	log "github.com/sirupsen/logrus"
)

var session *r.Session

func InitializeRethinkDB() {
	var err error
	var dbTimeout time.Duration = 100

	log.Info("Connecting to rethinkdb...")

	rdbAddress := fmt.Sprintf("%s:%s", os.Getenv("RDB_HOST"), os.Getenv("RDB_PORT"))
	log.Info(rdbAddress)

	for {
		if session, err = r.Connect(r.ConnectOpts{
			Address:  rdbAddress,
			Database: "test",
		}); err != nil {
			log.Error("Database connection denied, attempting to reconnect after ", dbTimeout, " miliseconds")
			time.Sleep(dbTimeout * time.Millisecond)
			dbTimeout *= 2

		} else {
			log.Info("Connection successful")

			break
		}

	}

	log.Info("Connection successful to rethinkdb")
}
