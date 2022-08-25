package logger

import (
	"fmt"
	"os"

	"github.com/arymaulanamalik/bobobox_test/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

const (
	logLevelTrace   = "TRACE"
	logLevelDebug   = "DEBUG"
	logLevelInfo    = "INFO"
	logLevelWarning = "WARNING"
	logLevelError   = "ERROR"
)

var logLevel = map[string]int{
	logLevelTrace:   5,
	logLevelDebug:   4,
	logLevelInfo:    3,
	logLevelWarning: 2,
	logLevelError:   1,
}

var activeLogLevel = config.LOG_LEVEL

func getActiveLogLevel() int {
	return logLevel[activeLogLevel]
}

func Warning(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[logLevelWarning] {

		strData := parseData(data)
		logrus.WithFields(logrus.Fields{
			"data":         strData,
			"service_name": config.SERVICE_NAME,
			"pod_name":     config.POD_NAME,
		}).Warning(message)

	}
}

func Debug(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[logLevelDebug] {
		strData := parseData(data)
		logrus.WithFields(logrus.Fields{
			"data":         strData,
			"service_name": config.SERVICE_NAME,
			"pod_name":     config.POD_NAME,
		}).Debug(message)
	}
}

func Info(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[logLevelInfo] {
		strData := parseData(data)
		logrus.WithFields(logrus.Fields{
			"data":         strData,
			"service_name": config.SERVICE_NAME,
			"pod_name":     config.POD_NAME,
		}).Info(message)
	}
}

func Sinfo(message string, data interface{}, span ddtrace.Span) {
	if getActiveLogLevel() >= logLevel[logLevelInfo] {
		strData := parseData(data)
		logrus.WithFields(logrus.Fields{
			"trace":        span,
			"data":         strData,
			"service_name": config.SERVICE_NAME,
			"pod_name":     config.POD_NAME,
		}).Info(message)
	}
}

func Error(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[logLevelError] {
		strData := parseData(data)
		logrus.WithFields(logrus.Fields{
			"data":         strData,
			"service_name": config.SERVICE_NAME,
			"pod_name":     config.POD_NAME,
		}).Error(message)
	}
}

func Serror(message string, data interface{}, span ddtrace.Span) {
	if getActiveLogLevel() >= logLevel[logLevelError] {
		strData := parseData(data)
		logrus.WithFields(logrus.Fields{
			"trace":        span,
			"data":         strData,
			"service_name": config.SERVICE_NAME,
			"pod_name":     config.POD_NAME,
		}).Error(message)
	}
}

func Fatal(message string, data interface{}) {
	Error(message, data)
	os.Exit(1)
}

func parseData(data interface{}) string {
	strData := ""
	if data != nil {
		strData = fmt.Sprintf("%v", data)
	}

	return strData
}
