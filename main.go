package main

import (
	"eniqilo_store/app"
	// "cats_social/controller"
	// "cats_social/exception"
	// "cats_social/repository"
	// "cats_social/service"
	"fmt"
	// "net/http"

	// "github.com/go-playground/validator/v10"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	// validate := validator.New()

	fmt.Println("Hello, World!")

	app.ConnectToPostgres()
	// db, err := app.ConnectToPostgres()
	if err != nil {
		fmt.Println(err)
	}



	// //user
	// userRepository := repository.NewUserRepository()
	// userService := service.NewUserService(userRepository, db, validate)
	// userController := controller.NewUserController(userService)

	// //cat
	// catRepository := repository.NewCatRepository()
	// catService := service.NewCatService(catRepository, db, validate)
	// catController := controller.NewCatController(catService)

	// router := app.NewRouter(userController, catController)
	// router.PanicHandler = exception.ErrorHandler
	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: router,
	// }
	// err = server.ListenAndServe()

}
