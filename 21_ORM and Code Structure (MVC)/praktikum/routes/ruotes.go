package route

import (
	controller "praktikum/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserIdController)
	e.POST("/users", controller.CreateUserIdController)
	e.DELETE("/users/:id", controller.DeleteUserIdController)
	e.PUT("/users/:id", controller.UpdateUserIdController)

	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookIdController)
	e.POST("/books", controller.CreateBookIdController)
	e.DELETE("/books/:id", controller.DeleteBookIdController)
	e.PUT("/books/:id", controller.UpdateBookIdController)

	e.GET("/blogs", controller.GetBlogsController)
	e.GET("/blogs/:id", controller.GetBlogIdController)
	e.POST("/blogs", controller.CreateBlogIdController)
	e.DELETE("/Blogs/:id", controller.DeleteBlogIdController)
	e.PUT("/blogs/:id", controller.UpdateBlogIdController)

	return e
}
