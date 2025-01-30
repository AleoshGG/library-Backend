package routes

import (
	"library-Backend/src/books/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	userRoutes := r.Group("/users") 
	{
		userRoutes.GET("/", controllers.GetUser)
	}
}