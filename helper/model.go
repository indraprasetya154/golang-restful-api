package helper

import (
	"github.com/indraprasetya154/golang-restful-api/model/domain"
	"github.com/indraprasetya154/golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
