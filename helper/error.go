package helper

import (
	"eniqilo_store/model/web"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func Unauthorized(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	webResponse := web.WebResponse{
		Message: "UNAUTHORIZED",
	}
	WriteToResponseBody(w, webResponse)
}

func BadRequest(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	webResponse := web.WebResponse{
		Message: message,
	}
	WriteToResponseBody(w, webResponse)
}
