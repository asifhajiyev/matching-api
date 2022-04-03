package model

type RestResponse struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	ErrorDetails interface{} `json:"errorDetails,omitempty"`
}

func BuildRestResponse(code int, message string, data interface{}, errorDetails interface{}) *RestResponse {
	return &RestResponse{
		Code:         code,
		Message:      message,
		Data:         data,
		ErrorDetails: errorDetails,
	}
}
