package main

import (
	"github.com/streadway/amqp"
	"log"
)

//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Fatalf("%s: %s", msg, err)
//	}
//}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal("failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"golang-queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal("Failed to register a consumer")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] waiting for message. To exit press Ctrl+C")
	<-forever
}
