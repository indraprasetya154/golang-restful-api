package service

import (
	"context"

	"github.com/indraprasetya154/golang-restful-api/model/web"
)

type CategoryServiceInterface interface {
	Create(ctx context.Context, recquest web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
