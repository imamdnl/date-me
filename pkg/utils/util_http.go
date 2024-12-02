package utils

import (
	"date-me/pkg/exceptions"
	"encoding/json"
	"net/http"
)

const (
	CODE_FAILED      = 1
	CODE_SUCCESS     = 0
	APPLICATION_JSON = "application/json"
	CONTENT_TYPE     = "Content-Type"
)

type BaseJsonResponse struct {
	Status DetailJsonHeadMessage `json:"status"`
}

type DetailJsonHeadMessage struct {
	Code               int      `json:"code"`
	Message            string   `json:"message"`
	DetailErrorMessage []string `json:"detailError,omitempty"`
}

type SuccessJsonResponse struct {
	BaseJsonResponse
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}

type FailedJsonResponse struct {
	DetailJsonHeadMessage
}

func NewSuccessJsonResponse(data interface{}, message string) SuccessJsonResponse {
	return SuccessJsonResponse{
		BaseJsonResponse: BaseJsonResponse{Status: DetailJsonHeadMessage{
			Code:    CODE_SUCCESS,
			Message: message,
		}},
		Data: data,
	}
}

func NewSuccessJsonResponseWithMeta(meta interface{}, data interface{}, message string) SuccessJsonResponse {
	return SuccessJsonResponse{
		BaseJsonResponse: BaseJsonResponse{Status: DetailJsonHeadMessage{
			Code:    CODE_SUCCESS,
			Message: message,
		}},
		Data: data,
		Meta: meta,
	}
}

func NewFailedJsonResponse(err error, message string) FailedJsonResponse {
	var messageErr string
	if err == nil {
		messageErr = message
	} else {
		messageErr = err.Error()
	}

	return FailedJsonResponse{
		DetailJsonHeadMessage{
			Code:    CODE_FAILED,
			Message: messageErr,
		}}
}

func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	jsonResponse := NewFailedJsonResponse(err, "Failed")
	response, _ := json.Marshal(jsonResponse)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	if _, err := w.Write(response); err == nil {
		return
	}

}

func RespondWithCustomError(w http.ResponseWriter, err *exceptions.CustomError) {
	var statusCode int
	switch err.Status {
	case exceptions.ERRDOMAIN:
		statusCode = http.StatusBadRequest
	case exceptions.ERRBUSSINESS:
		statusCode = http.StatusBadRequest
	case exceptions.ERRPRECONDITION:
		statusCode = http.StatusBadRequest
	case exceptions.ERRSYSTEM:
		statusCode = http.StatusInternalServerError
	case exceptions.ERRNOTFOUND:
		statusCode = http.StatusNotFound
	case exceptions.ERRREPOSITORY:
		statusCode = http.StatusBadRequest
	case exceptions.ERRUNKNOWN:
		statusCode = http.StatusBadRequest
	case exceptions.ERRAUTHORIZED:
		statusCode = http.StatusUnauthorized
	case exceptions.ERRFORBIDDEN:
		statusCode = http.StatusUnauthorized
	}
	jsonResponse := NewFailedJsonResponse(err.Errors, "Failed")
	response, _ := json.Marshal(jsonResponse)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	if _, err := w.Write(response); err == nil {
		return
	}
}

func RespondWithSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	successResponse := NewSuccessJsonResponse(data, "Success")
	response, _ := json.Marshal(successResponse)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	if _, err := w.Write(response); err == nil {
		return
	}
}

func ResponseWithSuccessMeta(w http.ResponseWriter, statusCode int, data interface{}, meta interface{}) {
	successResponse := NewSuccessJsonResponseWithMeta(meta, data, "Success")
	response, _ := json.Marshal(successResponse)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	if _, err := w.Write(response); err == nil {
		return
	}
}
