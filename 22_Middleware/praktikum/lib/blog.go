package database

import (
	"praktikum/config"
	"praktikum/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetBlogs() (interface{}, error) {
	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func GetBlog(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var blog models.Blog
	// if err := usecase.GetBlog(); err != nil{}
	if err := config.DB.First(&blog, id).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

func CreateBlog(c echo.Context) (interface{}, error) {
	blog := models.Blog{}
	c.Bind(&blog)
	if err := config.DB.Find(&models.User{}, blog.UserID).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Save(&blog).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

func DeleteBlog(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var blog models.Blog
	// var user models.User
	if err := config.DB.First(&blog, id).Error; err != nil {
		return nil, err
	}
	// config.DB.Model(&user).Where("id = ?", id).Association("Blogs").Clear()
	if err := config.DB.Delete(&blog, id).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

func UpdateBlog(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var blog models.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		return nil, err
	}
	c.Bind(&blog)
	if err := config.DB.Model(&blog).Where("id = ?", id).Updates(models.Blog{
		UserID: blog.UserID,
		Konten: blog.Konten,
		Judul:  blog.Judul,
	}).Error; err != nil {
		return nil, err
	}

	return blog, nil
}
