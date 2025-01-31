package controllers

import (
	"fmt"
	"library-Backend/src/books/aplication"
	"library-Backend/src/books/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetBookByTitleController struct {
	app *aplication.GetBookByTitle
}

func NewGetBookByTitleController() *GetBookByTitleController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewGetBookByTitle(mysql)
	return &GetBookByTitleController{app: app}
}

func (gbt_c *GetBookByTitleController) GetBookByTitle (c *gin.Context) {
	title := c.Param("title")
	fmt.Println(title)

	res := gbt_c.app.Run(title)

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