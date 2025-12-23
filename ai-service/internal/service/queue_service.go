package service

import (
	"ai-service/internal/event"
	"ai-service/internal/model"
	openai "ai-service/internal/openAi"


	"encoding/json"
	"fmt"
)

func ProcessMessage(body []byte, testService *TestService) {
	var req model.TestRequest
	if err := json.Unmarshal(body, &req); err != nil {
		fmt.Println(" Xabar parse qilinishida xatolik:", err)
		return
	}

questions, err := openai.ProcessPrompt(req.Prompt, req.Count)
if err != nil {
    // AI bilan ishlashdagi xatolikni unsuccess_queue ga yuborish
    if pubErr := event.PublishUnsuccess(testService.RabbitCh, req.TelegramID, err.Error()); pubErr != nil {
        fmt.Println("Unsuccess queue ga yuborishda xatolik:", pubErr)
    }
    return
}


	err = testService.CreateTest(req.TelegramID, questions)
	if err != nil {
		fmt.Println(" TestService xatosi:", err)
		return
	}

	
}
