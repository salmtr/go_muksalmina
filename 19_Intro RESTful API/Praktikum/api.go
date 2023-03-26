package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"user_id"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func GetBooksController(c echo.Context) error {
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts")

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var books []Book
	json.Unmarshal(responseData, &books)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to get Book",
		Status:  "OK",
		Data:    books,
	})

}

func GetByIdBooksController(c echo.Context) error {
	id := c.Param("id")
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var book Book
	json.Unmarshal(responseData, &book)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to get Book",
		Status:  "OK",
		Data:    book,
	})
}

func CreateBooksController(c echo.Context) error {
	book := new(Book)
	c.Bind(&book)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to create Book",
		Status:  "OK",
		Data:    book,
	})
}

func DeleteBooksController(c echo.Context) error {
	id := c.Param("id")
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var book Book
	json.Unmarshal(responseData, &book)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to delete Book",
		Status:  "OK",
		Data:    book,
	})
}

func main() {
	e := echo.New()
	e.GET("/books", GetBooksController)
	e.GET("/books/:id", GetByIdBooksController)
	e.POST("/books", CreateBooksController)
	e.DELETE("/books/:id", DeleteBooksController)
	e.Logger.Fatal(e.Start(":8080"))
}
