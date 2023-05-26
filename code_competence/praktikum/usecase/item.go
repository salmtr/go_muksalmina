package usecase

import (
	"code_competence/model"
	"code_competence/model/payload"
	"code_competence/repository/database"

	"github.com/google/uuid"
)

func CreateItem(req payload.CreateItemRequest) (*model.Item, error) {
	id, _ := uuid.NewRandom()
	item := &model.Item{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	item, err := database.CreateItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func GetItemByName(name string) ([]model.Item, error) {
	item, err := database.GetItemByName(name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func GetAllItems() ([]model.Item, error) {
	items, err := database.GetAllItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetItemByID(id string) (*model.Item, error) {
	item, err := database.GetItemByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func UpdateItemByID(id string, item model.Item) (*model.Item, error) {
	items, err := database.GetItemByID(id)
	if err != nil {
		return nil, err
	}
	items.Name = item.Name
	items.Description = item.Description
	items.Price = item.Price
	items.Stock = item.Stock
	items.CategoryID = item.CategoryID
	newItem, err := database.UpdateItem(items)
	if err != nil {
		return nil, err
	}
	return newItem, nil

}

func DeleteItemByID(id string) error {
	items, err := database.GetItemByID(id)
	if err != nil {
		return err
	}

	err = database.DeleteItem(items)
	if err != nil {
		return err
	}
	return nil
}
