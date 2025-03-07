package controllers

import (
	"library-Backend/src/books/aplication/useCases"
	"library-Backend/src/books/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllBooksController struct {
	app *aplication.GetAllBooks
}

func NewGetAllBooksController() *GetAllBooksController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewGetAllBooks(mysql)
	return &GetAllBooksController{app: app}
}

func (gb_c *GetAllBooksController) GetAllBooks(c *gin.Context) {
	res := gb_c.app.Run()

	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error": "No se consigió resultados",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/books/",
		},
		"data": res,
	})
}