package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)

	log.WithFields(log.Fields{
		"test": "01",
		"size": "big",
	}).Info("some info information")

}
