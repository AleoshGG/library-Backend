package main

import (
	booksInfra "library-Backend/src/books/infrastructure"
	rBooks "library-Backend/src/books/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	booksInfra.GoDependences()
	// Crear el router
	r := gin.Default()
	r.Use(cors.Default())


	// Registrar las rutas
	rBooks.RegisterRoutes(r)

	r.Run() // Sirve y escucha peticiones en 0.0.0.0:8080
}
