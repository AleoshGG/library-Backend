package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
  // Cargar las variables de entorno
  godotenv.Load()

  // Conexión con RABBITMQ
  conn, err := amqp.Dial(os.Getenv("URL"))
  failOnError(err, "Failed to connect to RabbitMQ")
  defer conn.Close()

  // Entrar a un canal
  ch, err := conn.Channel()
  failOnError(err, "Failed to open a channel")
  defer ch.Close()

  // Declaración al exchange (intercambiador) al cual éste consumidor se suscribira mediante una cola
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

  // Declaramos la cola a la cual estaremos suscritos
  q, err := ch.QueueDeclare(
    "", // name
    false,   // durable
    false,   // delete when unused
    false,   // exclusive
    false,   // no-wait
    nil,     // arguments
  )
  failOnError(err, "Failed to declare a queue")
  
  // (Data Binding) enlace de la cola con el exchange (intercambiador)
  err = ch.QueueBind(
    q.Name, // queue name
    "",     // routing key
    "logs", // exchange
    false,
    nil,
  )
  failOnError(err, "Failed to bind a queue")

  // Declaración de éste consumidor, que se suscribe a una cola
  msgs, err := ch.Consume(
    q.Name, // queue
    "",     // consumer
    true,   // auto-ack
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // args
  )
  failOnError(err, "Failed to register a consumer")

  // Tratamiento de un mensaje
  var forever chan struct{}

  go func() {
    for d := range msgs {
            log.Printf(" [x] %s", d.Body)
    }
  }()

  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}