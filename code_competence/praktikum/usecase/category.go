package usecase

import (
	"code_competence/model"
	"code_competence/model/payload"
	"code_competence/repository/database"
)

func CreateCategory(req payload.CategoryRequest) (*payload.CategoryResponse, error) {
	newCategory := &model.Category{
		Name: req.Name,
	}
	category, err := database.CreateCategory(newCategory)
	if err != nil {
		return nil, err
	}
	resp := payload.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return &resp, nil
}

func GetAllItemByCategoryID(id string) ([]model.Category, error) {
	category, err := database.GetAllItemByCategoryID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
