package core

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

//LoadEnv loads the file ".env" with the environment variables
func LoadEnv() {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	log.Info("File .env loaded correctly")
	log.Info("Rdb host and port:", os.Getenv("RDB_PORT"))
}
