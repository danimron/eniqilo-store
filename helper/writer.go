package helper

import (
	"encoding/json"
	"eniqilo_store/model/web"
	"eniqilo_store/pkg/errorwrapper"
	"net/http"
)

func Write(w http.ResponseWriter, data interface{}, err error) {
	if err == nil {
		err = errorwrapper.New(errorwrapper.StatusOK, nil, "")
	}

	errWrap := errorwrapper.CastToErrorWrapper(err)

	// in case error not nil but not wrapped
	if errWrap == nil {
		err = errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
		errWrap = errorwrapper.CastToErrorWrapper(err)
	}

	// write header
	w.Header().Set("Content-Type", "application/json")

	// write http status code
	w.WriteHeader(errWrap.HttpStatus)

	_ = json.NewEncoder(w).Encode(web.WebResponse{
		Message: errWrap.Message,
		Data:    data,
	})
}
