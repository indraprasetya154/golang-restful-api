package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/go-playground/validator/v10"
	"github.com/indraprasetya154/golang-restful-api/controller"
	"github.com/indraprasetya154/golang-restful-api/helper"
	"github.com/indraprasetya154/golang-restful-api/middleware"
	"github.com/indraprasetya154/golang-restful-api/model/domain"
	"github.com/indraprasetya154/golang-restful-api/repository"
	"github.com/indraprasetya154/golang-restful-api/routes"
	"github.com/indraprasetya154/golang-restful-api/service"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := routes.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCateogry(db *sql.DB) {
	db.Exec("TRUNCATE TABLE categories")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCateogry(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"category name test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "category name test", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCateogry(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad request", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCateogry(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "category name test",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"category name updated test"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "category name updated test", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCateogry(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "category name test",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad request", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCateogry(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "category name test",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCateogry(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/categories/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not found", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCateogry(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "category name test",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCateogry(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/categories/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not found", responseBody["status"])
}

func TestListCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCateogry(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "category name test",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/categories", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var categories = responseBody["data"].([]interface{})

	categoryResponse := categories[0].(map[string]interface{})

	assert.Equal(t, category.Id, int(categoryResponse["id"].(float64)))
	assert.Equal(t, category.Name, categoryResponse["name"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCateogry(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"category name test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
}
