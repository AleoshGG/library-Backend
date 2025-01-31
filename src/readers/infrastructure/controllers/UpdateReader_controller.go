package controllers

import (
	"library-Backend/src/readers/aplication"
	"library-Backend/src/readers/domain"
	"library-Backend/src/readers/infrastructure"
	"library-Backend/src/readers/infrastructure/controllers/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateReaderController struct {
	app *aplication.UpdateReader
}

func NewUpdateReaderController() *UpdateReaderController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewUpdateReader(mysql)
	return &UpdateReaderController{app: app}
}

func (ur_c *UpdateReaderController) UpdateReader(c *gin.Context) {
	id := c.Param("id")
	var newReader domain.Reader

	id_reader, _ := strconv.ParseInt(id, 10, 64)
	if err := c.ShouldBindJSON(&newReader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos: " + err.Error(),
		})
		return
	}

	if err := validators.CheckReader(newReader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	rowsAffected, _ := ur_c.app.Run(int(id_reader), newReader)

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "No se pudo actualizar el lector ",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": gin.H{
			"type": "reader",
			"id_reader": id_reader,
			"attributes": gin.H{
				"first_name": newReader.First_name,
				"last_name": newReader.Last_name,
				"email": newReader.Email,
				"phone_number": newReader.Phone_number,
			},
		},
	})
}