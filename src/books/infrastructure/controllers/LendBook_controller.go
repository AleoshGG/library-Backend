package controllers

import (
	"fmt"
	aplication "library-Backend/src/books/aplication/useCases"
	"library-Backend/src/books/domain"
	"library-Backend/src/books/infrastructure"
	"library-Backend/src/books/infrastructure/controllers/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LendBookController struct {
	app *aplication.UpdateBook
}

func NewLendBookController() *LendBookController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewUpdateBook(mysql)
	return &LendBookController{app: app}
}

func (ub_c *LendBookController) LendBook(c *gin.Context) {
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
	
	if err := validators.CheckBook(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	fmt.Println(book.Show())

	rowsAffected, _ := ub_c.app.Run(int(id_book), book)

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo actualizar el libro: No se entontró la referencia o ocurrió algo más",
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