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

type LendBookController struct {
	app *aplication.LendBook
	service *services.NotifyOfLendEvent
}

func NewLendBookController() *LendBookController {
	mysql := infrastructure.GetMySQL()
	rabbit := infrastructure.GetRabbitMQ()
	app := aplication.NewLendBook(mysql)
	service := services.NewNotifyOfLend(rabbit)
	return &LendBookController{app: app, service: service}
}

func (ub_c *LendBookController) LendBook(c *gin.Context) {
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

	// Notificar de prestamo
	ub_c.service.Run(int(loan.Id_reader), loan.Return_date)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}