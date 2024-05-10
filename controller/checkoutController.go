package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CheckoutController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	// FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
