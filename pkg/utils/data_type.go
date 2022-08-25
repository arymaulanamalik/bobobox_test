package utils

import (
	"encoding/json"
	"reflect"
)

// AnyTypeToString convert any type value to string
func AnyTypeToString(value interface{}) (valueString string) {
	switch valueType := value.(type) {
	case []byte:
		valueString = string(valueType)
	default:
		valueByte, _ := json.Marshal(value)
		valueString = string(valueByte)
	}
	return valueString
}

// ValueInArray is function for check value in list array
func ValueInArray(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		valueStored := reflect.ValueOf(array)

		for i := 0; i < valueStored.Len(); i++ {
			if reflect.DeepEqual(value, valueStored.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}
