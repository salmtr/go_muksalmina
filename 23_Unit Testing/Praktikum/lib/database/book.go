package database

import (
	"praktikum/config"
	"praktikum/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func CreateBook(c echo.Context) (interface{}, error) {
	book := models.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	c.Bind(&book)
	if err := config.DB.Model(&book).Where("id = ?", id).Updates(models.Book{
		Penulis:  book.Penulis,
		Penerbit: book.Penerbit,
		Judul:    book.Judul,
	}).Error; err != nil {
		return nil, err
	}

	return book, nil
}
