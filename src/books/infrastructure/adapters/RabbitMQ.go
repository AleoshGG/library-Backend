package adapters

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Notify struct {
	Id_reader   int
	Return_date string
}

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ() *RabbitMQ{
	conn, err := amqp.Dial(os.Getenv("URL_RABBIT"))
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return &RabbitMQ{conn: conn, ch: ch}  
}

func (r *RabbitMQ) NotifyOfLend(id_reader int, return_date string) {
	var notify Notify
	notify.Id_reader = id_reader
	notify.Return_date = return_date
	payload, err := json.Marshal(notify)
	failOnError(err, "Error al serializar Loan a JSON")
	r.prepareToMessage(payload)
}

func (r *RabbitMQ) NotifyOfReturn(id_reader int) {
	var notify Notify
	notify.Id_reader = id_reader
	notify.Return_date = "null"
	payload, err := json.Marshal(notify)
	failOnError(err, "Error al serializar Loan a JSON")
	r.prepareToMessage(payload)
}

func (r *RabbitMQ) prepareToMessage(body []byte) {
	// Declaración del exchange (intercambiador):
	r.ch.ExchangeDeclare(
		"exnotificacion",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	  
	r.ch.PublishWithContext(ctx,
		"exnotificacion",     // exchange
		"notificacion", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "application/json",
		  Body:        body,
		})
	//log.Printf(" [x] Sent %s\n", body)
}


func failOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
}