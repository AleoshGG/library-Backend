package controllers

import (
	"fmt"
	"library-Backend/src/books/aplication"
	"library-Backend/src/books/domain"
	"library-Backend/src/books/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateBookController struct {
	app *aplication.UpdateBook
}

func NewUpdateBookController() *UpdateBookController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewUpdateBook(mysql)
	return &UpdateBookController{app: app}
}

func (ub_c *UpdateBookController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book domain.Book

	id_book, _ := strconv.ParseInt(id, 10, 64)
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}
	
	fmt.Println(book.Show())

	rowsAffected, err := ub_c.app.Run(int(id_book), book)

	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo actualizar el libro " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": gin.H{
			"type": "book",
			"id_book": id_book,
			"attributes": gin.H{
				"title": book.Title,
				"date_publication": book.Date_publication,
				"editorial": book.Editorial,
				"amount": book.Amount,
			},
		},
	})
}