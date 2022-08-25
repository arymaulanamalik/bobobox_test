package utils

const (
	ErrorExpectedParameter           = "Request is invalid, please provide the expected parameter(s)"
	ErrorMinimalOneParameterToUpdate = "Request is invalid, please provide at least one expected parameter to update"
	ErrorValueShouldBePositive       = "Request is invalid, value should not be negative"
	ErrorInvalidDate                 = "Request is invalid, please provide the valid date"
	ErrorInvalidDateRange            = "Request is invalid, please provide the valid date range"
	ErrorParseRequestPayload         = "Failed to parse request payload"
	ErrorFormFileNoData              = "http: no such file"
	ErrorFormFileNotMultipartForm    = "request Content-Type isn't multipart/form-data"
)
