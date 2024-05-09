package app

import (
	"eniqilo_store/controller"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(staffController controller.StaffController, productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/v1/staff/register", staffController.Register)
	router.POST("/v1/staff/login", staffController.Login)

	// Product Route
	router.POST("/v1/product", productController.Create)
	router.PUT("/v1/product/:id", productController.Update)
	router.DELETE("/v1/product/:id", productController.Delete)

	return router
}

func printText(text string) {
	fmt.Println(text)
}
