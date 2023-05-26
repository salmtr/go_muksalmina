package database

import (
	"code_competence/config"
	"code_competence/model"
)

func CreateCategory(category *model.Category) (*model.Category, error) {
	err := config.DB.Create(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func GetAllItemByCategoryID(id string) ([]model.Category, error) {
	var category []model.Category
	err := config.DB.Preload("Item").Where("id = ?", id).Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
