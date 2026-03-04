package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY missing")
	}

	model := os.Getenv("GEMINI_MODEL")
	if model == "" {
		model = "gemini-1.5-flash"
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent", model)

	reqBody, err := json.Marshal(map[string]any{
		"contents": []any{
			map[string]any{
				"role": "user",
				"parts": []any{map[string]any{"text": fullPrompt}},
			},
		},
		"generationConfig": map[string]any{
			"responseMimeType": "application/json",
			"temperature":      0.2,
		},
	})
	if err != nil {
		return nil, err
	}

	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")


	req.Header.Set("x-goog-api-key", apiKey)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Debug: qayerda turib qolayotganini bilish uchun
	fmt.Println("[gemini] sending request to:", url, "model:", model)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("[gemini] status:", resp.Status)
	// agar error bo'lsa body’ni ko'rsatadi
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("gemini error: status=%s body=%s", resp.Status, string(respBytes))
	}

	// candidates[0].content.parts[0].text
	var parsed struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := json.Unmarshal(respBytes, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse gemini response: %w; raw=%s", err, string(respBytes))
	}

	if len(parsed.Candidates) == 0 || len(parsed.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no candidates/content: raw=%s", string(respBytes))
	}

	text := parsed.Candidates[0].Content.Parts[0].Text
	if text == "" {
		return nil, fmt.Errorf("empty text: raw=%s", string(respBytes))
	}

	var questions []Question
	if err := json.Unmarshal([]byte(text), &questions); err != nil {
		return nil, fmt.Errorf("AI returned invalid JSON: %w; raw=%s", err, text)
	}

	return questions, nil
}