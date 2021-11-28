package repository

import (
	"context"
	"database/sql"

	"github.com/indraprasetya154/golang-restful-api/model/domain"
)

type CategoryRepositoryInterface interface {
	Save(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx sql.Tx, category domain.Category) (domain.Category, error)
	FindAll(ctx context.Context, tx sql.Tx, category domain.Category) []domain.Category
}
