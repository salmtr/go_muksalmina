package route

import (
	"praktikum/constants"
	controller "praktikum/controllers"
	m "praktikum/middlewares"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	m.LoggerMiddleware(e)
	e.Pre(middleware.RemoveTrailingSlash())

	// Route / to handler function
	u := e.Group("/users")
	u.GET("", controller.GetUsersController, middleware.JWT([]byte(constants.SECRET_JWT)))
	u.GET("/:id", controller.GetUserController, middleware.JWT([]byte(constants.SECRET_JWT)))
	u.POST("", controller.CreateUserController)
	u.DELETE("/:id", controller.DeleteUserController, middleware.JWT([]byte(constants.SECRET_JWT)))
	u.PUT("/:id", controller.UpdateUserController, middleware.JWT([]byte(constants.SECRET_JWT)))
	u.POST("/login", controller.LoginUserController)

	b := e.Group("/books", middleware.JWT([]byte(constants.SECRET_JWT)))
	b.GET("", controller.GetBooksController)
	b.GET("/:id", controller.GetBookController)
	b.POST("", controller.CreateBookController)
	b.DELETE("/:id", controller.DeleteBookController)
	b.PUT("/:id", controller.UpdateBookController)

	bl := e.Group("/blogs")
	bl.GET("", controller.GetBlogsController)
	bl.GET("/:id", controller.GetBlogController)
	bl.POST("", controller.CreateBlogController)
	bl.DELETE("/:id", controller.DeleteBlogController)
	bl.PUT("/:id", controller.UpdateBlogController)

	return e
}
