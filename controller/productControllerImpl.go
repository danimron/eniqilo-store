package controller

import (
	"eniqilo_store/helper"
	"eniqilo_store/model/web"
	"eniqilo_store/pkg/errorwrapper"
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
	err := helper.NewReadFromRequestBody(r, &productCreateRequest)
	if err != nil {
		err = errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
		helper.Write(w, nil, err)
		return
	}
	productResponse, err := controller.ProductService.Create(r.Context(), productCreateRequest)
	helper.Write(w, productResponse, err)
}

func (controller *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		err = errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
		helper.Write(w, nil, err)
		return
	}

	productCreateRequest := web.ProductCreateReq{}
	productCreateRequest.Id = id

	err = helper.NewReadFromRequestBody(r, &productCreateRequest)
	if err != nil {
		err = errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
		helper.Write(w, nil, err)
		return
	}
	err = controller.ProductService.Update(r.Context(), productCreateRequest)
	helper.Write(w, nil, err)
}

func (controller *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		err = errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
		helper.Write(w, nil, err)
		return
	}

	err = controller.ProductService.Delete(r.Context(), id)
	helper.Write(w, nil, err)
}
