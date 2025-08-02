package service

import (
	"errors"
	"practice/gocast/todo-cli/server/model"
	"practice/gocast/todo-cli/server/store"
)

type CategoryService struct {
	categoryStore store.CategoryStore
}

func NewCategoryService(categoryStore store.CategoryStore) *CategoryService {
	return &CategoryService{categoryStore: categoryStore}
}

func (categoryService *CategoryService) AddCategory(title string, userId int) (*model.Category, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	category := &model.Category{
		Title:  title,
		UserID: userId,
	}
	err := categoryService.categoryStore.Create(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (categoryService *CategoryService) ListCategories(userId int) ([]*model.Category, error) {
	return categoryService.categoryStore.GetByUser(userId)
}
