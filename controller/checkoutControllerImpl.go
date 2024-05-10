package controller

import (
	"eniqilo_store/helper"
	"eniqilo_store/model/web"
	"eniqilo_store/pkg/errorwrapper"
	"eniqilo_store/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CheckoutControllerImpl struct {
	CheckoutService service.CheckoutService
}

func NewCheckoutController(checkoutService service.CheckoutService) CheckoutController {
	return CheckoutControllerImpl{
		CheckoutService: checkoutService,
	}
}
func (c CheckoutControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := web.CheckoutReq{}
	err := helper.NewReadFromRequestBody(r,&req)
	if err != nil{
		err = errorwrapper.New(errorwrapper.StatusInternalServerError, err,"")
		helper.Write(w,nil,err)
		return
	}
	err = c.CheckoutService.Create(r.Context(),req)
	
	helper.Write(w,nil,err)
}
