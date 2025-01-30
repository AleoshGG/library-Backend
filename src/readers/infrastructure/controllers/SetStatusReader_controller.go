package controllers

import (
	"library-Backend/src/readers/aplication"
	"library-Backend/src/readers/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SetStatusReaderController struct {
	app *aplication.SetStatusReader
}

func NewSetStatusReaderController() *SetStatusReaderController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewSetStatusReader(mysql)
	return &SetStatusReaderController{app: app}
}

func (sse_c *SetStatusReaderController) SetStatusReader(c *gin.Context) {
	id := c.Param("id")
	var status struct {
		Account_status string
	}

	id_reader, _ := strconv.ParseInt(id, 10, 64)
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos: " + err.Error(),
		})
		return
	}

	if status.Account_status != "active" && status.Account_status != "suspended"{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos, tienes que usar estos dos estados: 'active' o 'suspended'",
		})
		return
	}

	rowsAffected, err := sse_c.app.Run(int(id_reader), status.Account_status)

	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "No se pudo actualizar el estado " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": gin.H{
			"type":    "reader",
			"id_reader": id_reader,
			"attributes": gin.H{
				"account_status": status.Account_status,
			},
		},
	})
}