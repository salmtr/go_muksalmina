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
	e := echo.New()package route

	import (
		"code_competence/constant"
		"code_competence/controller"
		"net/http"
	
		"github.com/go-playground/validator"
		"gorm.io/gorm"
	
		"github.com/labstack/echo/v4"
		"github.com/labstack/echo/v4/middleware"
	)
	
	type customValidator struct {
		validator *validator.Validate
	}
	
	func (cv *customValidator) Validate(i interface{}) error {
		if err := cv.validator.Struct(i); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return nil
	}
	
	func NewRoute(e *echo.Echo, db *gorm.DB) {
	
		e.Validator = &customValidator{validator: validator.New()}
	
		// user route
		e.POST("/register", controller.RegisterUserController)
		e.POST("/login", controller.LoginUserController)
	
		// item route
		items := e.Group("/items", middleware.JWT([]byte(constant.SECRET_JWT)))
		items.POST("", controller.CreateItemController)
		items.GET("", controller.GetAllItemsController)
		items.GET("/id/:id", controller.GetItemByIDController)
		items.PUT("/id/:id", controller.UpdateItemByIDController)
		items.DELETE("/id/:id", controller.DeleteItemByIDController)
		items.GET("/name", controller.GetItemByItemNameController)
	
		// category route
		category := e.Group("/category", middleware.JWT([]byte(constant.SECRET_JWT)))
		category.POST("", controller.CreateCategoryController)
		category.GET("/:category_id", controller.GetAllByCategoryIDController)
	}
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
