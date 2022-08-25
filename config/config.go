package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	NAMESPACE                   string
	DOMAIN                      string
	TYPE                        string
	SERVICE_NAME                string
	API_PORT                    string
	LOG_LEVEL                   string
	STATE                       string
	POD_NAME                    string
	APP_SECRET_KEY              string
	MAX_HTTP_RETRY              int
	MYSQL_HOST                  string
	MYSQL_PORT                  int
	MYSQL_USER                  string
	MYSQL_PASSWORD              string
	MYSQL_DATABASE_NAME         string
	DD_AGENT_HOST               string
	DD_AGENT_PORT               string
	DD_VERSION                  string
	DD_AGENT_DEBUG              bool
	ENABLE_DD_PROFILER          bool
	ENABLE_DD_TRACE             bool
	TIME_ZONE                   string
	AUTH_SERVICE_API_URL        string
	PMS_SERVICE_API_URL         string
	USER_SERVICE_API_URL        string
	TUKANGPOS_SERVICE_API_URL   string
	TUKANGPOS_SERVICE_API_KEY   string
	SHARE_STAY_DYNAMIC_URL      string
	S3_URL                      string
	SAS_INSURANCE_API_IS_ACTIVE bool
	SAS_INSURANCE_API_URL       string
	SAS_INSURANCE_API_KEY       string
	SAS_TRACKER_API_URL         string
	SAS_TRACKER_API_KEY         string

	err error

	DateFormat         string = "2006-01-02"
	DateTimeFormat     string = "2006-01-02 15:04:05"
	FullDateTimeFormat string = "2006-01-02T15:04:05Z"
)

func init() {

	godotenv.Load()

	NAMESPACE = getEnv("NAMESPACE", "staging")
	DOMAIN = getEnv("DOMAIN", "sas")
	SERVICE_NAME = getEnv("SERVICE_NAME", "stay")
	TYPE = getEnv("TYPE", "api")
	POD_NAME = getEnv("POD_NAME", "stay")
	API_PORT = getEnv("API_PORT", "1212")
	LOG_LEVEL = getEnv("LOG_LEVEL", "DEBUG")
	STATE = getEnv("STATE", "DEVELOPMENT")
	APP_SECRET_KEY = getEnv("APP_SECRET_KEY", "3128e6f5-a00c-bc4a-67cc-2ac8be9aa453")
	MAX_HTTP_RETRY, _ = strconv.Atoi(getEnv("MAX_HTTP_RETRY", "2"))
	AUTH_SERVICE_API_URL = getEnv("AUTH_SERVICE_API_URL", "http://auth-service-api.staging.svc.cluster.local")
	PMS_SERVICE_API_URL = getEnv("PMS_SERVICE_API_URL", "http://pms-gateway-v2-api.staging.svc.cluster.local")
	USER_SERVICE_API_URL = getEnv("USER_SERVICE_API_URL", "http://sas-user-api.staging.svc.cluster.local")
	TUKANGPOS_SERVICE_API_URL = getEnv("TUKANGPOS_SERVICE_API_URL", "http://sas-tukangpos-api.staging.svc.cluster.local")
	TUKANGPOS_SERVICE_API_KEY = getEnv("TUKANGPOS_SERVICE_API_KEY", "a617ab76-f3d-f3f-726-e27-58701c14c383")
	SAS_INSURANCE_API_IS_ACTIVE, _ = strconv.ParseBool(getEnv("SAS_INSURANCE_API_IS_ACTIVE", "false"))
	SAS_INSURANCE_API_URL = getEnv("SAS_INSURANCE_API_URL", "")
	SAS_INSURANCE_API_KEY = getEnv("SAS_INSURANCE_API_KEY", "")
	SAS_TRACKER_API_URL = getEnv("SAS_TRACKER_API_URL", "")
	SAS_TRACKER_API_KEY = getEnv("SAS_TRACKER_API_KEY", "")
	// MYSQL database
	MYSQL_HOST = getEnv("MYSQL_HOST", "stag-pms-pms-m-0.stag-pms-pms-m.staging.svc.cluster.local")
	MYSQL_PORT, _ = strconv.Atoi(getEnv("MYSQL_PORT", "3306"))
	MYSQL_USER = getEnv("MYSQL_USER", "root")
	MYSQL_PASSWORD = getEnv("MYSQL_PASSWORD", "DCr2R9KPoP9U")
	MYSQL_DATABASE_NAME = getEnv("MYSQL_DATABASE_NAME", "pms")
	// Datadog Agent
	DD_AGENT_HOST = getEnv("DD_AGENT_HOST", "datadog-agent.platform.svc.cluster.local")
	DD_AGENT_PORT = getEnv("DD_AGENT_PORT", "812600")
	DD_VERSION = getEnv("DD_VERSION", "dev")
	DD_AGENT_DEBUG, _ = strconv.ParseBool(getEnv("DD_AGENT_DEBUG", "false"))
	ENABLE_DD_TRACE, err = strconv.ParseBool(getEnv("ENABLE_DD_TRACE", "false"))
	if err != nil {
		log.Fatal("FAILED TO PARSE ENABLE_DD_TRACE VALUE")
	}
	ENABLE_DD_PROFILER, err = strconv.ParseBool(getEnv("ENABLE_DD_PROFILER", "false"))
	if err != nil {
		log.Fatal("FAILED TO PARSE ENABLE_DD_PROFILER VALUE")
	}
	// Dynamic URL
	SHARE_STAY_DYNAMIC_URL = getEnv("SHARE_STAY_DYNAMIC_URL", "")
	// S3
	S3_URL = getEnv("S3_URL", "")

}

func getEnv(envKey string, defaultValue string) string {
	if value := os.Getenv(envKey); value != "" {
		return value
	}
	return defaultValue
}
