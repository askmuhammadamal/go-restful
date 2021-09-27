package main

import (
	"github.com/askmuhammadamal/go-restful-api/app"
	"github.com/askmuhammadamal/go-restful-api/controller"
	"github.com/askmuhammadamal/go-restful-api/helper"
	"github.com/askmuhammadamal/go-restful-api/middleware"
	"github.com/askmuhammadamal/go-restful-api/repository"
	"github.com/askmuhammadamal/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:        "localhost:3000",
		Handler:     middleware.NewAuthMiddleware(router),
		ReadTimeout: 60,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
