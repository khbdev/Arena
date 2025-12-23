package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"notifaction-service/internal/config"
	"notifaction-service/internal/model"
	"notifaction-service/internal/service"
	"sync"
)


func StartConsumer(queueName string, workerCount int) {
	ch := config.RabbitChannel()

	msgs, err := ch.Consume(
		queueName,
		"",  
		true, 
		false, 
		false, 
		false, 
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
			fmt.Printf("Worker %d started\n", workerID)

			for d := range msgs {
				var req model.TestRequest
				if err := json.Unmarshal(d.Body, &req); err != nil {
					log.Printf("Worker %d: JSON parse error: %v\n", workerID, err)
					continue
				}

			
				service.SendTestNotification(req.TelegramID, req.TestID)

			
			}
		}(i + 1)
	}

	wg.Wait()
}
