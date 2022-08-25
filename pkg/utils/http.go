package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/arymaulanamalik/bobobox_test/config"
	"github.com/arymaulanamalik/bobobox_test/model"
	"github.com/arymaulanamalik/bobobox_test/pkg/logger"
	"github.com/arymaulanamalik/bobobox_test/pkg/response"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Client is an HTTP client
var Client = &http.Client{
	Timeout: time.Second * time.Duration(120),
}

// DoHTTPRequest Do sends an HTTP request and returns an HTTP response
func DoHTTPRequest(request *http.Request) (response *http.Response, err error) {

	maxHTTPRetry := config.MAX_HTTP_RETRY

	for retry := 0; retry < maxHTTPRetry; retry++ {
		response, err = Client.Do(request)
		if err != nil {
			if retry == maxHTTPRetry-1 {
				return response, fmt.Errorf("Max HTTP retry exceeded. %s", err.Error())
			}
			time.Sleep(100 * time.Millisecond)
			continue
		}
		break
	}

	return response, nil
}

func SendHTTPRequest(ctx context.Context, method, url string, apiKey string, payload interface{}) (
	response response.Response, error model.Error) {

	span, _ := tracer.StartSpanFromContext(ctx, "SendPostHTTPRequest")
	defer span.Finish()

	// Prepare payload
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		error.Message = "Failed to prepare payload"
		error.Error = err
		logger.Serror(error.Message, error, span)
		return response, error
	}
	bytePayloadBuffer := bytes.NewBuffer(bytePayload)

	// Prepare request
	req, err := http.NewRequest(method, url, bytePayloadBuffer)
	if err != nil {
		error.Message = "Failed to prepare new request"
		error.Error = err
		logger.Serror(error.Message, error, span)
		return response, error
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Do HTTP request
	res, err := DoHTTPRequest(req)
	if err != nil {
		error.Message = "Failed in sending request"
		error.Error = err
		logger.Serror(error.Message, error, span)
		return response, error
	}
	defer res.Body.Close()

	// Process request response
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		error.Message = "Failed in reading response"
		error.Error = err
		logger.Serror(error.Message, error, span)
		return response, error
	}

	// Process response
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		error.Message = "Failed to decode response"
		error.Error = err
		logger.Serror(error.Message, error, span)
		return response, error
	}

	// Check Status Code
	if res.StatusCode != http.StatusOK {
		error.Code = fmt.Sprint(res.StatusCode)
		error.Message = response.Message
		error.Error = errors.New(error.Message)
		logger.Serror(error.Message, error, span)
		return response, error
	}

	return response, error
}
