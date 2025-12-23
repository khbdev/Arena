package handler

import (

	"log"
	"result-service/internal/config"
	"result-service/internal/service"
	"sync"

	"github.com/streadway/amqp"
)

const WorkerCount = 5 

func StartConsumer(rmq *config.RabbitMQ) {
	msgs, err := rmq.Channel.Consume(
		config.ResultQueue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	wg := &sync.WaitGroup{}
	msgChan := make(chan amqp.Delivery)


	for i := 0; i < WorkerCount; i++ {
		wg.Add(1)
		go worker(i, msgChan, wg)
	}

	
	go func() {
		for msg := range msgs {
			msgChan <- msg
		}
		close(msgChan)
	}()

	wg.Wait()
	log.Println("All workers finished")
}

func worker(id int, msgs <-chan amqp.Delivery, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range msgs {
		service.CalculateTestResult(msg.Body)
		msg.Ack(false)
	}
}

