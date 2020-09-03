package main

import (
	cor "go-api/core"
	utl "go-api/utils"

	"os"
)

func main() {
	cor.LoadEnv()
	utl.InitializeLogger()

	s := cor.Server{}
	s.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))
	s.Run(os.Getenv("PORT"))
}
