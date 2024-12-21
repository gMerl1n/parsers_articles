package errors

import (
	"encoding/json"
	"net/http"
)

type basicError struct {
	Success        bool   `json:"success"`
	Err            string `json:"error"`
	Cause          string `json:"cause"`
	httpStatusCode int
}

func (e *basicError) Error() string {
	return e.Err + ": " + e.Cause
}

func (e *basicError) SetCause(cause string) *basicError {
	e.Cause = cause

	return e
}

func (e *basicError) SetSuccess() *basicError {
	e.Success = true

	return e
}

func setCause(e *basicError, cause string) *basicError {
	return &basicError{
		Err:            e.Err,
		httpStatusCode: e.httpStatusCode,
		Cause:          cause,
	}
}

func new(err string, httpStatusCode int) *basicError {
	return &basicError{
		Err:            err,
		httpStatusCode: httpStatusCode,
	}
}

func SendHttpError(response http.ResponseWriter, errorMsg error) {

	responseJSON, err := json.Marshal(errorMsg)
	if err != nil {
		http.Error(response, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(responseJSON)

}
