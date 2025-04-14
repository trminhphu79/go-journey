package main

import (
	"app/framework"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Application init...")
	framework.Start()
}
