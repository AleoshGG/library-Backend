package controllers

import (
	"library-Backend/src/books/aplication/useCases"
	"library-Backend/src/books/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetBookByIdController struct {
	app *aplication.GetBookById
}

func NewGetBookByIdController() *GetBookByIdController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewGetBookById(mysql)
	return &GetBookByIdController{app: app}
}

func (gbi_c *GetBookByIdController) GetBookById(c *gin.Context) {
	id := c.Param("id")
	
	id_book, _ := strconv.ParseInt(id, 10, 64)

	res := gbi_c.app.Run(int(id_book))

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