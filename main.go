package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/indraprasetya154/golang-restful-api/app"
	"github.com/indraprasetya154/golang-restful-api/controller"
	"github.com/indraprasetya154/golang-restful-api/exception"
	"github.com/indraprasetya154/golang-restful-api/helper"
	"github.com/indraprasetya154/golang-restful-api/middleware"
	"github.com/indraprasetya154/golang-restful-api/repository"
	"github.com/indraprasetya154/golang-restful-api/service"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/categories", categoryController.FindAll)
	router.GET("/categories/:categoryId", categoryController.FindById)
	router.POST("/categories", categoryController.Create)
	router.PUT("/categories/:categoryId", categoryController.Update)
	router.DELETE("/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
