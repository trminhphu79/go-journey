package main

import (
	"app/startup"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Application init...")
	startup.Start()
}
