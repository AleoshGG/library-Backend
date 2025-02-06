package controllers

import (
	"fmt"
	"library-Backend/src/books/aplication"
	"library-Backend/src/books/infrastructure"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	mu        sync.Mutex
	clients   []chan string       // Lista de clientes conectados
	notifChan = make(chan string) // Canal para enviar nuevas notificaciones
)

func addClient() chan string {
	mu.Lock()
	defer mu.Unlock()
	client := make(chan string, 1)
	clients = append(clients, client)
	return client
}

func removeClient(client chan string) {
	mu.Lock()
	defer mu.Unlock()
	for i, c := range clients {
		if c == client {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}

func notifyClients(message string) {
	mu.Lock()
	defer mu.Unlock()
	for _, client := range clients {
		client <- message // Enviar la nueva notificación a cada cliente
	}
}

type Handler struct {
	app *aplication.GetAllBooks
}

func NewHandler() *Handler {
	mysql := infrastructure.GetMySQL()
	app := aplication.NewGetAllBooks(mysql)
	go monitorBookChanges(app)
	return &Handler{app: app}
}

func (gb_c *Handler) GetAllBooks(c *gin.Context) {
	clientChan := addClient() // Registrar al nuevo cliente
	defer func() {
		// Eliminar al cliente cuando termine la conexión
		removeClient(clientChan)
		close(clientChan)
	}()

	c.Header("Content-Type", "application/json")
	c.Header("Transfer-Encoding", "chunked")

	select {
	case msg := <-clientChan:
		c.JSON(http.StatusOK, gin.H{"notification": msg}) // Notificación
	case <-time.After(30 * time.Second): // Timeout de 30s
		c.JSON(http.StatusNoContent, nil)
	}
}

func monitorBookChanges(app *aplication.GetAllBooks) {
	// Este ciclo verifica los cambios en la base de datos
	for {
		time.Sleep(3 * time.Second) // Revisar cada 3s
		books := app.Run() // Obtener todos los libros de la base de datos
		notifyClients(fmt.Sprintf("Hay %d libros ahora", len(books))) // Notificar a los clientes
	}
}
