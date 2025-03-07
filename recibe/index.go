package main

import (
	"encoding/json"
	"log"
	"recibed/entities"
	"recibed/src"

	"github.com/joho/godotenv"
)

func main() {
  // Cargar las variables de entorno
  godotenv.Load()
  rabbit := src.NewRabbitMQ()
  
  // Tratamiento de un mensaje
  msgs := rabbit.GetMessages()
  var forever chan struct{}

  go func() {
    for d := range msgs {
        var loan entities.Loan
        err := json.Unmarshal(d.Body, &loan)
        if err != nil {
            log.Printf("Error al decodificar el mensaje: %s", err)
            continue
        }
        log.Printf(" [x] Prestamo recibido: id_reader=%d, id_book=%d, return_date=%s", loan.Id_reader, loan.Id_book, loan.Return_date)
        
        if loan.Return_date != "0000-00-00" {
          src.FetchAPI(loan, "lend")
        } else {
		  src.FetchAPI(loan, "return")
        } 
    }
}()

  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}


