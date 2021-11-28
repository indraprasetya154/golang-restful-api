package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/indraprasetya154/golang-restful-api/helper"
	"github.com/indraprasetya154/golang-restful-api/model/domain"
)

type CategoryRepository struct {
}

func (repository *CategoryRepository) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := `INSERT INTO categories (name) VALUES (?)`
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepository) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := `UPDATE categories SET name = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepository) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := `DELETE FROM categories WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepository) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := `SELECT id, name FROM categories WHERE id = ?`
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := `SELECT id, name FROM categories`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
