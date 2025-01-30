package controllers

import (
	"fmt"
	"library-Backend/src/readers/aplication"
	"library-Backend/src/readers/domain"
	"library-Backend/src/readers/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateReaderController struct {
	app *aplication.CreateReader
}

func NewCreateReaderController() *CreateReaderController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewCreateReader(mysql)
	return &CreateReaderController{app: app}
}

func (cr_c *CreateReaderController) CreateReader(c *gin.Context) {
	var reader domain.Reader
	reader.Account_status = "active"
	
	if err := c.ShouldBindJSON(&reader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}
	
	fmt.Println(reader.Show())
	
	id, err := cr_c.app.Run(reader)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo registrar al lector " + err.Error(),
		})
		return
	}
		
	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "book",
			"id_book": id,
			"attributes": gin.H{
				"first_name": reader.First_name,
				"last_name": reader.Last_name,
				"email": reader.Email,
				"phone_number": reader.Phone_number,
				"account_status": reader.Account_status,
			},
		},
	})
}