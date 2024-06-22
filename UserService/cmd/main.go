package main

import (
	"user_service/internal/app"
	"user_service/logger/lrus"
)

func main() {
	log := lrus.SetupLogger()
	application := app.New(log)
	go application.MustRun()
}
