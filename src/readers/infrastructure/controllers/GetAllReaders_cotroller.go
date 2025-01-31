package controllers

import (
	"library-Backend/src/readers/aplication"
	"library-Backend/src/readers/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllReadersController struct {
	app *aplication.GetAllReaders
}

func NewGetAllReadersController() *GetAllReadersController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewGetAllReaders(mysql)
	return &GetAllReadersController{app: app}
}

func (gar_c *GetAllReadersController) GetReaderByName(c *gin.Context) {
	readers := gar_c.app.Run()

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