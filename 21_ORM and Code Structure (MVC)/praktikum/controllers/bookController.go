package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"

	"github.com/labstack/echo/v4"
)

// Get all book controller
func GetBookController(c echo.Context) error {
	books, err := database.GetBook()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all book",
		Data:    books,
	})
}

// get book by id
func GetBookIdController(c echo.Context) error {
	book, err := database.GetBook(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"books":   book,
	})
}

// update book by id
func UpdateBookIdController(c echo.Context) error {
	book, err := database.UpdateBook(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})

}

// create new book
func CreateBookIdController(c echo.Context) error {
	book, err := database.CreateBook(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookIdController(c echo.Context) error {
	book, err := database.DeleteBook(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
		"book":    book,
	})
}
