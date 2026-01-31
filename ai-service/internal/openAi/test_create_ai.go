package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)



type Option struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Question struct {
	ID       string   `json:"id"`
	Question string   `json:"question"`
	Options  []Option `json:"options"`
	Correct  string   `json:"correct"`
}



// ProcessPrompt prompt bo‘yicha savollar generatsiya qiladi
// Agar AI noto‘g‘ri javob yoki limit tugasa, error qaytaradi
func ProcessPrompt(prompt string, count int) ([]Question, error) {
	fullPrompt := fmt.Sprintf(`
Generate %d questions about "%s" in JSON.
Each question must have:
- "id": unique string
- "question": text
- "options": array of objects with "id": "A"-"D" and "text": option text
- "correct": the id of the correct option

Return ONLY valid JSON array.
Do NOT explain anything.
Do NOT use markdown.
`, count, prompt)

	// Request body
	body, err := json.Marshal(map[string]interface{}{
		"model": "gpt-4.1",
		"input": fullPrompt,
		"store": true,
	})
	if err != nil {
		return nil, err
	}

	// API key
	token := os.Getenv("OPENAI_API_KEY")
	if token == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY missing")
	}

	// HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/responses", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	

	var parsed map[string]interface{}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	outputArr, ok := parsed["output"].([]interface{})
	if !ok || len(outputArr) == 0 {
		return nil, fmt.Errorf("no output from OpenAI")
	}

	outputObj, ok := outputArr[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid output object")
	}

	contentArr, ok := outputObj["content"].([]interface{})
	if !ok || len(contentArr) == 0 {
		return nil, fmt.Errorf("no content in output")
	}

	contentObj, ok := contentArr[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid content object")
	}

	text, ok := contentObj["text"].(string)
	if !ok || text == "" {
		return nil, fmt.Errorf("text not found in content")
	}

	var questions []Question
	if err := json.Unmarshal([]byte(text), &questions); err != nil {
		return nil, fmt.Errorf("AI returned invalid JSON: %w", err)
	}

	return questions, nil
}
