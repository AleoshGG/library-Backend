package controllers

import (
	"library-Backend/src/books/aplication/services"
	aplication "library-Backend/src/books/aplication/useCases"
	"library-Backend/src/books/domain"
	"library-Backend/src/books/infrastructure"
	"library-Backend/src/books/infrastructure/controllers/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReturnBookController struct {
	app *aplication.ReturnBook
	service *services.NotifyOfReturnEvent
}

func NewReturnBookController() *ReturnBookController {
	mysql := infrastructure.GetMySQL()
	rabbit := infrastructure.GetRabbitMQ()
	app := aplication.NewReturnBook(mysql)
	service := services.NewNotifyOfReturnEvent(rabbit)
	return &ReturnBookController{app: app, service: service}
}

func (ub_c *ReturnBookController) ReturnBook(c *gin.Context) {
	var loan domain.Loan

	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}
	
	if err := validators.CheckLoan(loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	rowsAffected, _ := ub_c.app.Run(int(loan.Id_book))

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo actualizar el libro: No se entontró la referencia o ocurrió algo más",
		})
		return
	}

	// Notificar de devuelto
	ub_c.service.Run(int(loan.Id_reader))

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}