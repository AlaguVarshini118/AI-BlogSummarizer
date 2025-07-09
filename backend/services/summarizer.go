package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	

	"github.com/joho/godotenv"
)

const cohereURL = "https://api.cohere.ai/v1/summarize"

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("❌ Error loading .env file")
	} else {
		fmt.Println("✅ .env loaded successfully")
	}
}

func SummarizeText(text string) (string, error) {
	apiKey := os.Getenv("COHERE_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("Cohere API key not set")
	}

	if len(text) < 250 {
		fmt.Println("⚠️ Text too short for Cohere, using local summarizer")
		return callLocalSummarizer(text)
	}

	reqBody := map[string]interface{}{
		"text":   text,
		"length": "medium",
		"format": "paragraph",
		"model":  "summarize-xlarge",
	}

	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", cohereURL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", err
	}

	if errMsg, exists := result["message"]; exists {
		return "", fmt.Errorf("Cohere error: %v", errMsg)
	}

	summary, ok := result["summary"].(string)
	if !ok || summary == "" {
		return "", fmt.Errorf("No summary returned by Cohere")
	}

	return summary, nil
}

func callLocalSummarizer(text string) (string, error) {
	payload := map[string]string{"text": text}
	body, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:5001/summarize", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("Local summarizer failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]string
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}

	summary := result["summary"]
	if summary == "" {
		return "", fmt.Errorf("Local summary empty")
	}

	return summary, nil
}
