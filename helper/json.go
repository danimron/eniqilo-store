package helper

import (
	"encoding/json"
	"eniqilo_store/pkg/errorwrapper"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body) // untuk decode json menjadi struct
	err := decoder.Decode(result)
	PanicIfError(err)
}

func NewReadFromRequestBody(r *http.Request, result interface{}) error {
	decoder := json.NewDecoder(r.Body) // untuk decode json menjadi struct
	err := decoder.Decode(result)
	if err != nil {
		return errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
	}
	return nil
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w) // untuk encode kembali menjadi json
	err := encoder.Encode(response)
	PanicIfError(err)
}
