package database

import (
	"praktikum/config"
	"praktikum/middlewares"
	"praktikum/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Preload("Blogs").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := config.DB.Preload("Blogs").First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	c.Bind(&user)
	if err := config.DB.Model(&user).Where("id = ?", id).Updates(models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
