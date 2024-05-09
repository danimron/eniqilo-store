package exception

import (
	"eniqilo_store/helper"
	"eniqilo_store/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if SomethingErrorHandler(w, r, err) {
		return
	}
	if validationErrors(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		// Code:    http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Data:    err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func SomethingErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(SomethingError)
	if ok {
		if exception.NotFoundError != "" {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusNotFound)
			webResponse := web.WebResponse{
				// Code:    http.StatusNotFound,
				Message: "NOT FOUND",
				Data:    exception.NotFoundError,
			}

			helper.WriteToResponseBody(writer, webResponse)
			return true
		} else if exception.PasswordError != "" {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusNotFound)
			webResponse := web.WebResponse{
				// Code:    http.StatusUnauthorized,
				Message: "INCORRECT PASSWORD",
				Data:    exception.PasswordError,
			}

			helper.WriteToResponseBody(writer, webResponse)
			return true
		} else {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusNotFound)
			webResponse := web.WebResponse{
				// Code:    http.StatusBadRequest,
				Message: "FOUND ERROR",
				Data:    exception.somethingError,
			}

			helper.WriteToResponseBody(writer, webResponse)
			return true
		}

	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			// Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}
