package controllers

import (
	"fmt"
	"library-Backend/src/books/aplication"
	"library-Backend/src/books/domain"
	"library-Backend/src/books/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookController struct {
	app *aplication.CreateBook
}

func NewCreateBookController() *CreateBookController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewCreateBook(mysql)
	return &CreateBookController{app: app}
}

func (cb_c *CreateBookController) AddBook(c *gin.Context) {
	var book domain.Book
	
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}
	
	fmt.Println(book.Show())
	
	id, err := cb_c.app.Run(book)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo guardar el libro " + err.Error(),
		})
		return
	}
		
	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "book",
			"id_book": id,
			"attributes": gin.H{
				"title": book.Title,
				"date_publication": book.Date_publication,
				"editorial": book.Editorial,
				"amount": book.Amount,
			},
		},
	})
}