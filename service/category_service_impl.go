package service

import (
	"context"
	"database/sql"
	"github.com/askmuhammadamal/go-restful-api/exception"
	"github.com/askmuhammadamal/go-restful-api/helper"
	"github.com/askmuhammadamal/go-restful-api/model/domain"
	"github.com/askmuhammadamal/go-restful-api/model/web"
	"github.com/askmuhammadamal/go-restful-api/repository"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (services *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := services.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := services.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = services.CategoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (services *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := services.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := services.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := services.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = services.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (services *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := services.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := services.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	services.CategoryRepository.Delete(ctx, tx, category)
}

func (services *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := services.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := services.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (services *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := services.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := services.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
