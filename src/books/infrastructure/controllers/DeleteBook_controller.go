package controllers

import (
	"library-Backend/src/books/aplication/useCases"
	"library-Backend/src/books/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteBookController struct {
	app *aplication.DeleteBook
}

func NewDeleteBookController() *DeleteBookController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewDeleteBook(mysql)
	return &DeleteBookController{app: app}
}

func (db_c *DeleteBookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	id_book, _ := strconv.ParseInt(id, 10, 64)

	rowsAffected, _ := db_c.app.Run(int(id_book))

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo actualizar el libro: No se entontró la referencia o ocurrió algo más",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Recurso eliminado",
	})
}