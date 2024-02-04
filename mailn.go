package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books = []Book{
	{ID: "1", Title: "Денискины рассказы", Author: &Author{Firstname: "Виктор", Lastname: "Драгунский"}},
	{ID: "2", Title: "Маленький принц", Author: &Author{Firstname: "Антуан", Lastname: "де Сент-Экзюпери"}},
}

func main() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.GET("/books", getBooks)
	router.GET("/book/:id", getBooksByID)
	router.POST("/book", postBook)
	router.PUT("/book/:id", updateBookByID)
	router.DELETE("/book/:id", deleteBookByID)
	router.Run("localhost:8080")
}

func updateBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, book := range books {
		if book.ID == id {
			var book Book
			if err := ctx.BindJSON(&book); err != nil {
				return
			}
			books[i] = book
			ctx.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, p := range books {
		if p.ID == id {
			books = append(books[:i], books[i+1:]...)
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func postBook(ctx *gin.Context) {
	var book Book

	if err := ctx.BindJSON(&book); err != nil {
		return
	}

	books = append(books, book)
	ctx.IndentedJSON(http.StatusCreated, book)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBooksByID(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, book := range books {
		if book.ID == id {
			ctx.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
