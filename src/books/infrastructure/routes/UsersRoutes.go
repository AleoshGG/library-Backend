package routes

import (
	"server01/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	userRoutes := r.Group("/users") 
	{
		userRoutes.GET("/", controllers.GetUser)
	}
}