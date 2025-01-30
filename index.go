package main

import (
	"library-Backend/src/books/infrastructure/routes"
	"library-Backend/src/core"

	"github.com/gin-gonic/gin"
)

func main() {
	core.GetDBPool()

	// Crear el router
	r := gin.Default()

	// Registrar las rutas
	routes.RegisterRoutes(r)

	r.Run() // Sirve y escucha peticiones en 0.0.0.0:8080
}
