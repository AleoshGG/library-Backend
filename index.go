package main

import (
	booksInfra "library-Backend/src/books/infrastructure"
	rBooks "library-Backend/src/books/infrastructure/routes"
	readerInfra "library-Backend/src/readers/infrastructure"
	rReaders "library-Backend/src/readers/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	booksInfra.GoMySQL()
	readerInfra.GoMySQL()
	// Crear el router
	r := gin.Default()
	r.Use(cors.Default())


	// Registrar las rutas
	rBooks.RegisterRoutes(r)
	rReaders.RegisterRoutes(r)

	r.Run() // Sirve y escucha peticiones en 0.0.0.0:8080
}
