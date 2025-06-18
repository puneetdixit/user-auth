package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	// Open or create the log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.SetOutput(os.Stdout)
		log.Warn("Failed to log to file, using default stderr")
	} else {
		log.SetOutput(file)
	}

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)
}
