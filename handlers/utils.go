package handlers

import (
	"api/contracts"
	"api/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	var httpErr httpError
	switch err.(type) {
	case errors.ValidationError:
		valErr := err.(errors.ValidationError)
		logrus.Errorf("validation failed: Error: %v, Details: %v", valErr.Code, err.(errors.ValidationError).Details)
		httpErr = httpError{
			GenericError: err.(errors.ValidationError).GenericError,
			StatusCode:   http.StatusBadRequest,
		}
	case errors.NotFoundError:
		logrus.Errorf("data not found")
		httpErr = httpError{
			GenericError: err.(errors.NotFoundError).GenericError,
			StatusCode:   http.StatusNotFound,
		}

	default:
		inErr := errors.NewInternalError(err)
		logrus.Errorf("unexpected error: %v", err)
		httpErr = httpError{
			GenericError: inErr.GenericError,
			StatusCode:   http.StatusInternalServerError,
		}
	}
	writeResponse(w, httpErr.StatusCode, &contracts.ErrorResponse{
		Success: false,
		Data:    nil,
		Errors:  []errors.GenericError{httpErr.GenericError},
	})
}

type httpError struct {
	errors.GenericError
	StatusCode int `json:"-"`
}
