package database

import (
	"code_competence/config"
	"code_competence/model"
)

func CreateUser(user *model.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
