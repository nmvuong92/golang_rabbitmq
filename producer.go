package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(err, "Failed to connect to RabbitMQ server")

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"golang-queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	failOnError(err, "Failed to declare a queue")

	body := "Golang is awesome - Keep moving forward!"

	err = ch.Publish(
		"",     //exchange
		q.Name, // route key
		false,  //mandatory
		false,  //immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	failOnError(err, "Failed to publishing a message")
	log.Printf("[x] congrats, sending message: %s\n", body)
}
