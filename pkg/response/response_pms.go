package response

import (
	"encoding/json"
	"net/http"

	"github.com/arymaulanamalik/bobobox_test/pkg/logger"
)

func SuccessPMS(res http.ResponseWriter, httpCode int, data interface{}, pagination interface{}) {
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(httpCode)

	response := SuccessResponsePMS{
		Status:  httpCode,
		Code:    0,
		Message: "success",
		Data:    data,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		logger.Fatal(err.Error(), nil)
	}
}
