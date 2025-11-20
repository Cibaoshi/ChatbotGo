package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type GeminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
	PromptFeedback struct {
		SafetyRatings []struct {
			Category    string `json:"category"`
			Probability string `json:"probability"`
		} `json:"safetyRatings"`
		BlockReason string `json:"blockReason"`
	} `json:"promptFeedback"`
}

func main() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY not found")
	}

	if len(os.Args) < 2 {
		log.Fatal("Enter your question in args")
	}
	userPrompt := os.Args[1]

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=" + apiKey

	req := GeminiRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{Text: userPrompt},
				},
			},
		},
	}

	body, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("marshal error: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("http error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		log.Fatalf("API error %d: %s", resp.StatusCode, responseBody)
	}

	var result GeminiResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("decode error: %v", err)
	}

	if len(result.Candidates) == 0 {
		if result.PromptFeedback.BlockReason != "" {
			fmt.Println("Blocked:", result.PromptFeedback.BlockReason)
		} else {
			fmt.Println("No candidates returned")
		}
		return
	}

	candidate := result.Candidates[0]
	if len(candidate.Content.Parts) == 0 {
		fmt.Println("Candidate has no parts")
		return
	}

	fmt.Println(candidate.Content.Parts[0].Text)
}
