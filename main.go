package main

import (
	"QRGenerator/app"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Application Started")
	app.ApiServer()
	defer log.Info("Application finish")
}
