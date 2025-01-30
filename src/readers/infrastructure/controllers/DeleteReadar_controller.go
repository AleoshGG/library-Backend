package controllers

import (
	"library-Backend/src/readers/aplication"
	"library-Backend/src/readers/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteReaderController struct {
	app *aplication.DeleteReader
}

func NewDeleteReaderController() *DeleteReaderController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewDeleteReader(mysql)
	return &DeleteReaderController{app: app}
}

func (dr_c *DeleteReaderController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	id_reader, _ := strconv.ParseInt(id, 10, 64)

	rowsAffected, err := dr_c.app.Run(int(id_reader))

	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "No se pudo eliminar el lector " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Recurso eliminado",
	})
}