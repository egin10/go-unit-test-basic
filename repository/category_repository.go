package repository

import "go-unit-test-basic/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
