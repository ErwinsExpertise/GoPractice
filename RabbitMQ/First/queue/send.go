package queue

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/streadway/amqp"
)

type Config struct {
	user     string
	password string
	host     string
	port     string
}

var conf Config
var conn *amqp.Connection

func init() {
	_, err := toml.Decode("config.toml", &conf)
	failOnError(err, "Failed decoding the configuration")
}

func Send(msg, qName string) {
	conn, err := amqp.Dial("amqp://" + conf.user + ":" + conf.password + "@" + conf.host + ":" + conf.port)
	failOnError(err, "Connection to RabbitMQ failed")

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		qName, //name
		false, //durable
		false, //delete when unused
		false, //exclusive
		false, //no-wait
		nil,   //arguments
	)

	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     //exchange
		q.Name, //routing key
		false,  //mandatory
		false,  //immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})

	log.Printf("[x] Sent %s", msg)

	failOnError(err, "Failed to publish a message")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("\nMessage: %s\nError: %s", msg, err)
	}
}
