package controllers

import (
	"library-Backend/src/readers/aplication"
	"library-Backend/src/readers/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetReaderByNameController struct {
	app *aplication.GetReaderByName
}

func NewGetReaderByNameController() *GetReaderByNameController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewGetReaderByName(mysql)
	return &GetReaderByNameController{app: app}
}

func (gbn_c *GetReaderByNameController) GetReaderByName(c *gin.Context) {
	name := c.Param("name")

	readers := gbn_c.app.Run(name)

	if readers == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error": "No se consigió resultados",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/readers/",
		},
		"data": readers,
	})
} 