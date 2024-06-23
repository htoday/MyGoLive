package dao

import "github.com/streadway/amqp"

var Ch *amqp.Channel
var Conn *amqp.Connection

func NewProducer(ch *amqp.Channel, roomId string, content string) error {
	routingKey := "chat." + roomId
	err := ch.Publish(
		"chat",     // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(content)},
	)
	return err
}

func NewConsumer(ch *amqp.Channel, roomId string) (<-chan amqp.Delivery, error) {
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}
	routingKey := "chat." + roomId
	err = ch.QueueBind(
		q.Name,     // queue name
		routingKey, // routing key
		"chat",     // exchange
		false,
		nil)
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
		return nil, err
	}

	return msgs, nil
}
