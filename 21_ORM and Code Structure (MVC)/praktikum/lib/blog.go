package database

import (
	"praktikum/config"
	"praktikum/models"
)

func GetBlog() (blogs []models.Blog, err error) {
	err = config.DB.Find(&blogs).Error

	if err != nil {
		return []models.Blog{}, err
	}

	return
}

func CreateBlog(blogs models.Blog) (models.Blog, error) {
	err := config.DB.Create(&blogs).Error

	if err != nil {
		return models.Blog{}, err
	}

	return blogs, nil
}

func GetBlogId(id any) (models.Blog, error) {
	blogs := models.Blog{}

	err := config.DB.Where("id = ?", id).First(&blogs).Error

	if err != nil {
		return models.Blog{}, err
	}

	return blogs, nil
}

func UpdateBlog(blogs models.Blog, id any) (models.Blog, error) {
	err := config.DB.Where("id = ?", id).Updates(&blogs).Error

	if err != nil {
		return models.Blog{}, err
	}

	return blogs, nil
}

func DeleteBlog(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Blog{}).Error

	if err != nil {
		return nil, err
	}

	return nil, nil
}
