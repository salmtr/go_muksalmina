package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid user ID",
		})
	}

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user by ID",
				"user":     user,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid user ID",
		})
	}

	index := -1
	for i, user := range users {
		if user.Id == id {
			index = i
			break
		}
	}

	if index >= 0 {
		users = append(users[:index], users[index+1:]...)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "success delete user",
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid user ID",
		})
	}

	var updatedUser User
	if err := c.Bind(&updatedUser); err != nil {
		return err
	}

	exists := false
	for i, user := range users {
		if user.Id == id {
			users[i] = updatedUser
			exists = true
			break
		}
	}

	if exists {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "success update user",
			"user":     updatedUser,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
