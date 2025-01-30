package routes

import (
	"library-Backend/src/readers/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	readersRouter := r.Group("/readers") 
	{
		readersRouter.POST("/newReader", controllers.NewCreateReaderController().CreateReader)
		readersRouter.GET("/q=:name", controllers.NewGetReaderByNameController().GetReaderByName)
	}
}