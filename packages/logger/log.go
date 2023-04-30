package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const ProgramName = "go-webhttp-backend"

func New() *log.Entry {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime: "@timestamp",
			log.FieldKeyMsg:  "message",
		},
	})
	logger.SetLevel(log.TraceLevel)
	logger.SetOutput(os.Stdout)
	return logger.WithField("PROGRAM", ProgramName)
}
