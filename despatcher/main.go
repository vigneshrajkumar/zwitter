package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
)

const tweetQueue string = "tweet"

func main() {
	log.Println("Desptcher up and running")

	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		tweetQueue, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Fatalln(err)
	}

	tweet := make(map[string]interface{}, 0)
	tweet["userID"] = 1
	tweet["contents"] = "Microservices getting up and running #boom"
	tweet["timestamp"] = time.Now().Unix()
	tweet["location"] = "Chennai, India"

	message, err := json.Marshal(tweet)
	if err != nil {
		log.Println(err)
	}

	err = ch.Publish("", /* sending to the default exchange*/
		queue.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: message})
	if err != nil {
		log.Println(err)
	}

	log.Println("sent sample message")
}
