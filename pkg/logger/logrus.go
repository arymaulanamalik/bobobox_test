package logger

import (
	"os"

	"github.com/arymaulanamalik/bobobox_test/config"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var (
	lognew = log.New()
	logger *log.Entry
)

func init() {

	lognew.Out = os.Stdout

	logger = lognew.WithFields(
		log.Fields{
			"service_name": config.SERVICE_NAME,
			"pod_name":     config.POD_NAME,
		},
	)

	if config.STATE == "DEVELOPMENT" {
		logrus.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: false,
		})
	} else {
		logrus.SetFormatter(&log.JSONFormatter{
			PrettyPrint: false,
		})
	}
}
