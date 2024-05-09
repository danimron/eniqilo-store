package main

import (
	"eniqilo_store/app"
	"eniqilo_store/controller"
	"eniqilo_store/exception"
	"eniqilo_store/repository"
	"eniqilo_store/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	validate := validator.New()

	fmt.Println("Hello, World!")

	app.ConnectToPostgres()
	db, err := app.ConnectToPostgres()
	if err != nil {
		fmt.Println(err)
	}

	//staff
	staffRepository := repository.NewStaffRepo()
	staffService := service.NewStaffService(staffRepository, db, validate)
	staffController := controller.NewStaffController(staffService)

	//product
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := app.NewRouter(staffController, productController)
	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err = server.ListenAndServe()

}
