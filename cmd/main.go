package main

import (
	"app/platform"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Application init...")
	platform.Start()
}
