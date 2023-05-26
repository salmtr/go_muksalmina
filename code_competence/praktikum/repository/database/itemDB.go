package database

import (
	"code_competence/config"
	"code_competence/model"
)

func CreateItem(item *model.Item) (*model.Item, error) {
	err := config.DB.Create(item).Error
	if err != nil {
		return nil, err
	}
	err = config.DB.Preload("Category").First(item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func GetItemByName(name string) ([]model.Item, error) {
	var items []model.Item

	err := config.DB.Where("name LIKE ? ", "%"+name+"%").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetAllItems() ([]model.Item, error) {
	var items []model.Item
	err := config.DB.Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetItemByID(id string) (*model.Item, error) {
	var item model.Item
	err := config.DB.Where("id = ?", id).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func UpdateItem(item *model.Item) (*model.Item, error) {
	err := config.DB.Save(item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func DeleteItem(item *model.Item) error {
	err := config.DB.Delete(&item).Error
	if err != nil {
		return err
	}
	return nil
}
