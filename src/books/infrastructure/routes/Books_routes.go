package routes

import (
	"library-Backend/src/books/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	userRoutes := r.Group("/books") 
	{
		userRoutes.POST("/newBook", controllers.NewCreateBookController().AddBook)
		userRoutes.GET("/", controllers.NewGetAllBooksController().GetAllBooks)
		userRoutes.GET("/:id", controllers.NewGetBookByIdController().GetBookById)
		userRoutes.GET("/q=:title", controllers.NewGetBookByTitleController().GetBookByTitle)
		userRoutes.PUT("/:id", controllers.NewUpdateBookController().UpdateBook)
	}
}