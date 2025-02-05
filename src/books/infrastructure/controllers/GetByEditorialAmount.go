package controllers

import (
	"library-Backend/src/books/aplication"
	"library-Backend/src/books/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetByEditorialAmountController struct {
	app *aplication.GetByEditorialAmount
}

func NewGetByEditorialAmount() *GetByEditorialAmountController {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewGetByEditorialAmount(mysql)
	return &GetByEditorialAmountController{app: app}
}

func (gbea_c *GetByEditorialAmountController) GetByEditorialAmount(c *gin.Context) {
	editorial := c.Query("editorial")
    amount := c.Query("amount") 
	
	minAmount, _ := strconv.ParseInt(amount, 10, 64)

	res := gbea_c.app.Run(editorial, int(minAmount))

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