package dto

// struct  ErrorResponse
type ErrorResponse struct {
	ResponseMeta
	Data   any `json:"data"`
	Errors any `json:"errors,omitempty"`
}

// struct DefaultErrorResponse
func DefaultErrorResponse() ErrorResponse {
	return DefaultErrorResponseWithMessage("")
}

// Struct DefaultErrorResponseWithMessage
func DefaultErrorResponseWithMessage(msg string) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      msg,
			ResponseTime: "",
		},
		Data: nil,
	}
}

// DefaultErrorInvalidDataWithMessage
func DefaultErrorInvalidDataWithMessage(msg string) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      "Form Invalid data.",
			ResponseTime: "",
		},
		Data: msg,
	}
}

// struct DefaultDataInvalidResponse
func DefaultDataInvalidResponse(validationErrors any) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			MessageTitle: "Oops, something went wrong.",
			Message:      "Data invalid.",
		},
		Errors: validationErrors,
	}
}

// struct DefaultBadRequestResponse
func DefaultBadRequestResponse() ErrorResponse {
	return DefaultErrorResponseWithMessage("Bad request")
}
