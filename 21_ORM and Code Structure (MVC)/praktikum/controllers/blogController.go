package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"

	"github.com/labstack/echo/v4"
)

// Get all Blog controller
func GetBlogController(c echo.Context) error {
	Blogs, err := database.GetBlogs()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all blogs",
		Data:    Blogs,
	})
}

// get Blog by id
func GetBlogIdController(c echo.Context) error {
	Blog, err := database.GetBlog(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog",
		"Data":    Blog,
	})
}

// update Blog by id
func UpdateBlogIdController(c echo.Context) error {
	Blog, err := database.UpdateBlog(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update Blog",
		"Data":    Blog,
	})

}

// create new Blog
func CreateBlogIdController(c echo.Context) error {
	Blog, err := database.CreateBlog(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new Blog",
		"Data":    Blog,
	})
}

// delete Blog by id
func DeleteBlogIdController(c echo.Context) error {
	Blog, err := database.DeleteBlog(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Blog",
		"Data":    Blog,
	})
}
