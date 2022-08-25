package response

type SuccessResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Code       int         `json:"code,omitempty"`
}

type SuccessResponsePMS struct {
	Status  int         `json:"status,omitempty"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Params  Params      `json:"params"`
}
type Params struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalItems int `json:"totalItems"`
}

type ErrorResponse struct {
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Pagination struct {
	Page       int `json:"page,omitnull"`
	DataInPage int `json:"data_in_page,omitnull"`
	TotalData  int `json:"total_data,omitnull"`
	TotalPage  int `json:"total_page,omitnull"`
}

type Response struct {
	Code       int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Pagination Pagination  `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
