package app

import (
	"github.com/indraprasetya154/golang-restful-api/controller"
	"github.com/indraprasetya154/golang-restful-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryControllerInterface) *httprouter.Router {
	router := httprouter.New()

	router.GET("/categories", categoryController.FindAll)
	router.GET("/categories/:categoryId", categoryController.FindById)
	router.POST("/categories", categoryController.Create)
	router.PUT("/categories/:categoryId", categoryController.Update)
	router.DELETE("/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
