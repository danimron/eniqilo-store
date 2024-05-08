package app

import (
	// "eniqilo_store/controller"
	// "eniqilo_store/middleware"
	"fmt"

	// "github.com/julienschmidt/httprouter"
)

// func NewRouter(userController controller.UserController, catController controller.CatController) *httprouter.Router {
// 	router := httprouter.New()

// 	router.POST("/v1/user/register", userController.Register)
// 	router.POST("/v1/user/login", userController.Login)

// 	router.POST("/v1/cat", middleware.VerifyToken(catController.Create))
// 	router.GET("/v1/cat", middleware.VerifyToken(catController.FindAll))
// 	router.DELETE("/v1/cat/:cat_id", middleware.VerifyToken(catController.Delete))
// 	router.PUT("/v1/cat/:cat_id", middleware.VerifyToken(catController.Update))
// 	// router.POST("/v1/user/register", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// 	printText("Register Route")
// 	// })
// 	// router.POST("/v1/user/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// 	printText("Login Route")
// 	// })
// 	// router.POST("/v1/cat/match", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// 	printText("Match Cat Route")
// 	// })
// 	// router.GET("/v1/cat/match", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// 	printText("Get Cat Match Route")
// 	// })
// 	// router.POST("/v1/cat/match/approve", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// 	printText("Approve Cat Match Route")
// 	// })
// 	// router.POST("/v1/cat/match/reject", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// 	printText("Reject Cat Match Route")
// 	// })
// 	// router.DELETE("/v1/cat/match/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// 	printText("Delete Cat Match Route")
// 	// })

// 	return router
// }

func printText(text string) {
	fmt.Println(text)
}
