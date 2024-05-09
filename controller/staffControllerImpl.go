package controller

import (
	"eniqilo_store/helper"
	"eniqilo_store/model/web"
	"eniqilo_store/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type StaffControllerImpl struct {
	StaffService service.StaffService
}

func NewStaffController(staffService service.StaffService) StaffController {
	return &StaffControllerImpl{
		StaffService: staffService,
	}
}

func (controller *StaffControllerImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	staffCreateRequest := web.StaffRegisterReq{}
	helper.ReadFromRequestBody(r, &staffCreateRequest)
	staffResponse, err := controller.StaffService.Register(r.Context(), staffCreateRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		webResponse := web.WebResponse{
			Message: err.Error(),
		}
		helper.WriteToResponseBody(w, webResponse)
	} else {
		w.WriteHeader(http.StatusCreated)
		webResponse := web.WebResponse{
			Message: "Staff registered successfully",
			Data:    staffResponse,
		}
		helper.WriteToResponseBody(w, webResponse)
	}
}

func (controller *StaffControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	staffLoginRequest := web.StaffLoginReq{}
	helper.ReadFromRequestBody(r, &staffLoginRequest)
	staffResponse, err := controller.StaffService.Login(r.Context(), staffLoginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		webResponse := web.WebResponse{
			Message: err.Error(),
		}
		helper.WriteToResponseBody(w, webResponse)
	} else {
		w.WriteHeader(http.StatusOK)
		webResponse := web.WebResponse{
			Message: "Staff logged successfully",
			Data:    staffResponse,
		}
		helper.WriteToResponseBody(w, webResponse)
	}
}
