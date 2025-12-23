package handler

import (
	"encoding/json"
	"log"
	"notifaction-service/internal/config"
	"notifaction-service/internal/model"
	"notifaction-service/internal/service"
	"sync"
)

func StartUserStaticConsumer(queueName string, workerCount int) {
	ch := config.RabbitChannel()

	msgs, err := ch.Consume(
		queueName,
		"",
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			defer wg.Done()
			log.Printf("Worker %d started\n", workerID)

			for d := range msgs {
				var result model.TestResult
				if err := json.Unmarshal(d.Body, &result); err != nil {
					log.Printf("Worker %d: JSON parse error: %v\n", workerID, err)
					continue
				}

			
				service.SendTestResult(result)
			}
		}(i + 1)
	}

	wg.Wait()
}
