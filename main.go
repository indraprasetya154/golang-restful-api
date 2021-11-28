package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	"github.com/go-playground/validator/v10"
	"github.com/indraprasetya154/golang-restful-api/controller"
	"github.com/indraprasetya154/golang-restful-api/database"
	"github.com/indraprasetya154/golang-restful-api/helper"
	"github.com/indraprasetya154/golang-restful-api/middleware"
	"github.com/indraprasetya154/golang-restful-api/repository"
	"github.com/indraprasetya154/golang-restful-api/routes"
	"github.com/indraprasetya154/golang-restful-api/service"
	"github.com/indraprasetya154/golang-restful-api/utils"
)

func main() {
	// init config
	utils.InitConfig()

	db := database.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := routes.NewRouter(categoryController)

	server := http.Server{
		Addr:    viper.GetString("APP_URL"),
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
