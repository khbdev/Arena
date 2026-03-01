package service

import (
	"ai-service/internal/domain"
	"ai-service/internal/event"
	"ai-service/internal/model"
	openai "ai-service/internal/openAi"

	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type MessageProcessor struct {
	testService domain.TestCreator
	RabbitCh *amqp.Channel
}


func NewMessageProcessor(testService domain.TestCreator, ch *amqp.Channel) *MessageProcessor {
	return &MessageProcessor{
		testService: testService,
		RabbitCh: ch,
	}
}


func (p *MessageProcessor) ProcessMessage(body []byte) {
	var req model.TestRequest

	if err := json.Unmarshal(body, &req); err != nil {
		fmt.Println("Xabar parse qilinishida xatolik:", err)
		return
	}

	questions, err := openai.ProcessPrompt(req.Prompt, req.Count)
	if err != nil {
		if pubErr := event.PublishUnsuccess(
			p.RabbitCh,
			req.TelegramID,
			err.Error(),
		); pubErr != nil {
			fmt.Println("Unsuccess queue ga yuborishda xatolik:", pubErr)
		}
		return
	}

	if err := p.testService.CreateTest(req.TelegramID, questions); err != nil {
		fmt.Println("TestService xatosi:", err)
		return
	}
}