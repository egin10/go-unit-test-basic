package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-unit-test-basic/entity"
	"go-unit-test-basic/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	//	Program Mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNilf(t, err, err.Error())
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	categoryResult, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNilf(t, categoryResult, "Success")
	assert.Equal(t, category.Id, categoryResult.Id)
	assert.Equal(t, category.Name, categoryResult.Name)
}
