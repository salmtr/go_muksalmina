package controllers

import (
	"net/http"
	"praktikum/lib/database"

	"github.com/labstack/echo"
)

func GetBlogsController(c echo.Context) error {
	blogs, err := database.GetBlogs()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"Blogs":   blogs,
	})
}

// get Blog by id
func GetBlogController(c echo.Context) error {
	blog, err := database.GetBlog(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"Blogs":   blog,
	})
}

// create new Blog
func CreateBlogController(c echo.Context) error {
	blog, err := database.CreateBlog(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// update Blog by id
func UpdateBlogController(c echo.Context) error {
	blog, err := database.UpdateBlog(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    blog,
	})

}

// delete Blog by id
func DeleteBlogController(c echo.Context) error {
	blog, err := database.DeleteBlog(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
		"blog":    blog,
	})
}
