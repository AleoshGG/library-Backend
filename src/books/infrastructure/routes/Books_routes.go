package routes

import (
	"library-Backend/src/books/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	bookRoutes := r.Group("/books") 
	{
		bookRoutes.POST("/", controllers.NewCreateBookController().AddBook)
		bookRoutes.GET("/", controllers.NewGetAllBooksController().GetAllBooks)
		bookRoutes.GET("/:id", controllers.NewGetBookByIdController().GetBookById)
		bookRoutes.GET("/q=:title", controllers.NewGetBookByTitleController().GetBookByTitle)
		bookRoutes.PUT("/lend/:id", controllers.NewLendBookController().LendBook)
		bookRoutes.PUT("/return/:id", controllers.NewReturnBookController().ReturnBook)
		bookRoutes.DELETE("/:id", controllers.NewDeleteBookController().DeleteBook)
		bookRoutes.GET("/search", controllers.NewGetByEditorialAmount().GetByEditorialAmount)
	}
}