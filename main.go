package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//book struct
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

//home struct
type home struct {
	Version     string `json:"version"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

//books data
var books = []book{
	{ID: "1", Title: "Go Beginner", Author: "Graham Katana", Quantity: 20},
	{ID: "2", Title: "PHP Beginner", Author: "Graham Katana", Quantity: 10},
	{ID: "3", Title: "Javascript Beginner", Author: "Graham Katana", Quantity: 12},
	{ID: "4", Title: "Dart for Flutter", Author: "Graham Katana", Quantity: 2},
	{ID: "5", Title: "Laravel", Author: "Graham Katana", Quantity: 0},
}

//get all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

//create a new book
func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

//find a book by id
func findBook(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)

}

//checkout book from library
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing an id query parameter"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book is not available currently"})
		return

	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

//return a book to libray
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing an id query parameter"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

//helper function to get by ID
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")

}

func loadHome(c *gin.Context) {
	var hero = []home{
		{
			Version:     "1.0.0",
			Title:       "Contoso Library",
			Author:      "Graham K Katana",
			Description: "Welcome to golang API",
		},
	}
	c.IndentedJSON(http.StatusOK, hero)

}

func main() {
	//gin router setup
	router := gin.Default()

	router.GET("/", loadHome)
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", findBook)
	router.PATCH("/books/checkout", checkoutBook)
	router.PATCH("/books/return", returnBook)

	router.Run("localhost:8080")

}
