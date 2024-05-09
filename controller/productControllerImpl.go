package controller

import (
	"eniqilo_store/helper"
	"eniqilo_store/model/web"
	"eniqilo_store/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productCreateRequest := web.ProductCreateReq{}
	helper.ReadFromRequestBody(r, &productCreateRequest)
	productResponse := controller.ProductService.Create(r.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Message: "a",
		Data:    productResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

// func (controller *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	CatGetParam := web.CatGetParam{}
// 	helper.NewGetCatParam(r, &CatGetParam)

// 	catResponse := controller.CatService.FindAll(r.Context(), &CatGetParam)
// 	webResponse := web.WebResponse{
// 		Message: "a",
// 		Data:    catResponse,
// 	}
// 	helper.WriteToResponseBody(w, webResponse)
// }

func (controller *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("id")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productCreateRequest := web.ProductCreateReq{}
	productCreateRequest.Id = id

	helper.ReadFromRequestBody(r, &productCreateRequest)
	controller.ProductService.Update(r.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Message: "a",
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("id")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	controller.ProductService.Delete(r.Context(), id)

	webResponse := web.WebResponse{
		Message: "a",
	}
	helper.WriteToResponseBody(w, webResponse)
}
