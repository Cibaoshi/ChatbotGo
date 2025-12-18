package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func sendToGemini(prompt, apiKey string) (string, error) {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=" + apiKey

	reqBody := GeminiRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{Text: prompt},
				},
			},
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("ошибка HTTP-запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ошибка API: HTTP статус %d. Ответ: %s", resp.StatusCode, responseBody)
	}

	var result GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("ошибка десериализации JSON: %v", err)
	}

	if len(result.Candidates) == 0 {
		if result.PromptFeedback.BlockReason != "" {
			return "", fmt.Errorf("модель заблокирована: %s", result.PromptFeedback.BlockReason)
		}
		return "", fmt.Errorf("модель не вернула ответ")
	}

	if len(result.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("ответ кандидата пуст")
	}

	return result.Candidates[0].Content.Parts[0].Text, nil
}
