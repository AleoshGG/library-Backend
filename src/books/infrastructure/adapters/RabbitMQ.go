package adapters

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {}

func NewRabbitMQ() *RabbitMQ{
	return &RabbitMQ{} 
}

func (r *RabbitMQ) NotifyOfLend() {
	// Conectar con nuestro host de RABBITMQ
	conn, err := amqp.Dial(os.Getenv("URL_RABBIT"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Entramos al canal
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declaración del exchange (intercambiador):
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	  
	// Creación de una publicación
	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(ctx,
		"logs",     // exchange
		"", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "text/plain",
		  Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	
	log.Printf(" [x] Sent %s\n", body)
}

func (r *RabbitMQ) NotifyOfReturn() {
	
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "Va a la misma cola, el exchange entrega a todos"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
  

func failOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
}