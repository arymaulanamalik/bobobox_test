package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ParseErrorJSONDecoder parse errors JSON decoder
func ParseErrorJSONDecoder(err error) error {
	var errJSONDecoder struct {
		Field  interface{}
		Offset interface{}
		Struct interface{}
		Type   interface{}
		Value  interface{}
	}

	// JSON encoding error
	val, _ := json.Marshal(err)

	// parses the JSON-encoded data and stores the result in the value pointed to struct errJSONDecoder.
	json.Unmarshal(val, &errJSONDecoder)

	msg := fmt.Sprintf("Incorrect data type format of '%v' field", errJSONDecoder.Field)

	if errJSONDecoder.Field == nil {
		msg = fmt.Sprintf("Incorrect data type format of fields")
	}

	return errors.New(msg)
}
