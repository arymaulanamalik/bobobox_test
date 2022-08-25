package utils

import (
	"time"

	"github.com/arymaulanamalik/bobobox_test/config"
)

var (
	TZ_ASIA_JAKARTA string = "Asia/Jakarta"
	TZ_UTC          string = "UTC"
)

var (
	DateFormat                 string = "2006-01-02"
	DateTimeFormat             string = "2006-01-02 15:04:05"
	FullDateTimeFormat         string = "2006-01-02T15:04:05Z"
	MySQLDatetimeFormat        string = "%Y-%m-%d %T"
	MySQLDateFormat            string = "%Y-%m-%d"
	ISO8601DateTimeFormat      string = "2006-01-02T15:04:05.999Z07:00"
	BCAVARequestFormatDatetime string = "02/01/2006 15:04:05"
	BcaDateTimeFormat          string = "2006-01-02 15:04:05 -0700"
)

func StrDateTimeToDatetime(strDateTime string) (time.Time, error) {
	dateTime, err := time.Parse(config.DateTimeFormat, strDateTime)
	if err != nil {
		return dateTime, err
	}
	return dateTime, nil
}

func MysqlStrDateTimeToDatetime(strDateTime string) (time.Time, error) {
	dateTime, err := time.Parse(config.FullDateTimeFormat, strDateTime)
	if err != nil {
		return dateTime, err
	}
	return dateTime, nil
}

func DateTimeToStrDateTime(dateTime time.Time) string {
	return dateTime.Format(config.DateTimeFormat)
}

func DateTimeToStrDate(dateTime time.Time) string {
	return dateTime.Format(config.DateFormat)
}

func StrDatetimeToDatetime(strDatetime string) (time.Time, error) {
	date, err := time.Parse(config.DateTimeFormat, strDatetime)
	if err != nil {
		return date, err
	}

	return date, nil
}

func StringMysqlDatetimeNow() string {

	return time.Now().Format(config.DateTimeFormat)

}

func CurrentTimeDateTime() time.Time {

	currentTimeStr := StringMysqlDatetimeNow()
	currentTime, _ := StrDateTimeToDatetime(currentTimeStr)
	return currentTime
}
