package events

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/events/eventhandlers"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"sync"
)

var channel *amqp091.Channel

const TwitterExchangeName = "twitterExchange"
const TwitterQueueName = "twitterQueue"
const TwitterRoutingKey = "twitterRoutingKey"

func PublishEventMessage(ctx context.Context, addr string, data interface{}) {
	bytesOfData, err := json.Marshal(data)
	if err != nil {
		log.Println("failed to marshal data")
		return
	}
	eventMessage := EventMessage{
		EventHandlerName: addr,
		Message:          string(bytesOfData),
	}
	bytesOfEventMessage, err := json.Marshal(eventMessage)
	if err != nil {
		log.Println("failed to marshal event message")
		return
	}

	err = channel.PublishWithContext(
		ctx,
		TwitterExchangeName,
		TwitterRoutingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        bytesOfEventMessage,
		})
	if err != nil {
		log.Println("failed to publish!")
	}
}

func ChannelDeclareQueue(ch *amqp091.Channel) error {
	_, err := ch.QueueDeclare(
		TwitterQueueName,
		true,
		false,
		false,
		false,
		nil)
	return err
}

func ChannelDeclareExchange(ch *amqp091.Channel) error {
	err := ch.ExchangeDeclare(
		TwitterExchangeName,
		amqp091.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil)
	return err
}

func ChannelBindQueue(ch *amqp091.Channel) error {
	err := ch.QueueBind(
		TwitterQueueName,
		TwitterRoutingKey,
		TwitterExchangeName,
		false,
		nil)
	return err
}

func ChannelConsume(ch *amqp091.Channel) (<-chan amqp091.Delivery, error) {
	return ch.Consume(
		TwitterQueueName,
		"",
		false,
		false,
		false,
		false,
		nil)
}

func GetEventHandlers() map[string]func(delivery amqp091.Delivery) {
	followEvtHandlerInstance := BaseEventHandler[models.FollowRequestModel]{
		ExecutorFunc: eventhandlers.FollowEventHandler,
	}
	unfollowEvtHandlerInstance := BaseEventHandler[models.FollowRequestModel]{
		ExecutorFunc: eventhandlers.UnfollowEventHandler,
	}

	return map[string]func(delivery amqp091.Delivery){
		eventhandlers.FollowEventHandlerName:   followEvtHandlerInstance.GetHandler(),
		eventhandlers.UnfollowEventHandlerName: unfollowEvtHandlerInstance.GetHandler(),
	}
}

func InitRabbitMq(runState *sync.WaitGroup) {
	failOnError := func(err error, msg string) {
		if err != nil {
			log.Println(fmt.Sprintf("%s : %s", msg, err))
			runState.Done()
		}
	}

	go func() {
		conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "failed to dial")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "failed initializing channel")
		if channel == nil {
			channel = ch
		}
		defer ch.Close()

		log.Println("channel initiated")

		err = ChannelDeclareQueue(ch)
		failOnError(err, "failed to declare queue")
		err = ChannelDeclareExchange(ch)
		failOnError(err, "failed to declare exchange")
		err = ChannelBindQueue(ch)
		failOnError(err, "failed to bind queue")

		eventHandlers := GetEventHandlers()
		for {
			messages, _ := ChannelConsume(ch)
			for msg := range messages {
				var evtMsg EventMessage
				err := json.Unmarshal(msg.Body, &evtMsg)
				if err != nil {
					log.Println("error json-serializing event message")
				}

				if _, ok := eventHandlers[evtMsg.EventHandlerName]; ok {
					log.Println(fmt.Sprintf("Event Received with Valid Handler Name: %s", evtMsg.EventHandlerName))
					eventHandlers[evtMsg.EventHandlerName](msg)
				} else {
					log.Println(fmt.Sprintf("No event handlers found w/ name: %s", evtMsg.EventHandlerName))
					err = msg.Ack(false)
					if err != nil {
						log.Println("error acknowledging event: " + err.Error())
					}
				}
			}
		}
	}()
}

type EventMessage struct {
	EventHandlerName string `json:"eventHandlerName"`
	Message          string `json:"message"`
}

type BaseEventHandler[T interface{}] struct {
	ExecutorFunc func(data T)
}

func (b *BaseEventHandler[T]) GetHandler() func(d amqp091.Delivery) {
	return func(d amqp091.Delivery) {
		data, err := ParseEventMessage[T](d)
		if err != nil {
			log.Println("error parsing amqp data!")

			err = d.Ack(false)
			if err != nil {
				log.Println("parse amqp data failed -> acknowledgement failed!")
			}
			return
		}
		b.ExecutorFunc(data)

		err = d.Ack(false)
		if err != nil {
			log.Println("event w/ valid event handler name ack failed")
		}
	}
}

func ParseEventMessage[T interface{}](d amqp091.Delivery) (T, error) {
	var eventMessage EventMessage
	err := json.Unmarshal(d.Body, &eventMessage)
	if err != nil {
		log.Println("error unmarshalling event message!")
		return *new(T), err
	}

	var data T
	err = json.Unmarshal([]byte(eventMessage.Message), &data)
	if err != nil {
		log.Println("error unmarshalling message")
		return *new(T), err
	}

	return data, nil
}
