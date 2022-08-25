package response

import (
	"encoding/json"
	"net/http"

	"github.com/arymaulanamalik/bobobox_test/pkg/logger"
)

func Success(res http.ResponseWriter, httpCode int, data interface{}, pagination interface{}) {
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(httpCode)

	response := SuccessResponse{
		Code:       httpCode,
		Data:       data,
		Pagination: pagination,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		logger.Fatal(err.Error(), nil)
	}
}

func Error(res http.ResponseWriter, httpCode int, message string) {
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(httpCode)

	response := ErrorResponse{
		Message: message,
		Code:    httpCode,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		logger.Fatal(err.Error(), nil)
	}
}
