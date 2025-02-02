package routes

import (
	"library-Backend/src/readers/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	readersRouter := r.Group("/readers") 
	{
		readersRouter.POST("/", controllers.NewCreateReaderController().CreateReader)
		readersRouter.GET("/q=:name", controllers.NewGetReaderByNameController().GetReaderByName)
		readersRouter.GET("/", controllers.NewGetAllReadersController().GetReaderByName)
		readersRouter.DELETE("/:id", controllers.NewDeleteReaderController().DeleteBook)
		readersRouter.PUT("/status/:id", controllers.NewSetStatusReaderController().SetStatusReader)
		readersRouter.PUT("/:id", controllers.NewUpdateReaderController().UpdateReader)
	}
}