package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	RepeatPenalty    float64   `json:"repeat_penalty"`
	RepeatLastN      int       `json:"repeat_last_n"`
	TopK             int       `json:"top_k"`
	TopP             float64   `json:"top_p"`
	MinP             float64   `json:"min_p"`
	XtcThreshold     float64   `json:"xtc-threshold"`
	Temperature      float64   `json:"temperature"`
	DryMultiplier    float64   `json:"dry_multiplier"`
	DryBase          float64   `json:"dry_base"`
	DryAllowedLength int       `json:"dry_allowed_length"`
	DryPenaltyLastN  int       `json:"dry_penalty_last_n"`
	Stream           bool      `json:"stream"`
}

func main() {
	// Parse command-line arguments
	modelName := flag.String("model", "gpt-3.5-turbo", "The name of the model to use")
	serverURL := flag.String("url", "http://localhost:8080/v1/chat/completion", "The server URL")
	flag.Parse()

	// Prepare the request payload
	requestBody := ChatRequest{
		Model: *modelName,
		Messages: []Message{
			{
				Role: "system",
				Content: "Translate the given text into idiomatic, simple, and accessible Vietnamese with natural southern Vietnamese semantics and idioms. " +
					"The translation should be straightforward enough for uneducated laypersons to understand, avoiding technical terms or specific Buddhist connotations. " +
					"Stay faithful to the original text by providing a verbatim 1:1 translation without paraphrasing, summarizing, or omitting any content. " +
					"Ensure that the translation flows cohesively while preserving cultural and spiritual connotations in a way that resonates with the target audience.",
			},
			{
				Role:    "user",
				Content: "Middle Discourses 141 The Analysis of the Truths 1.1So I have heard. 1.2At one time the Buddha was staying near Varanasi, in the deer park at Isipatana. 1.3There the Buddha addressed the mendicants, 1.4“Mendicants!” 1.5“Venerable sir,” they replied. 1.6The Buddha said this: 2.1“Near Varanasi, in the deer park at Isipatana, the Realized One, the perfected one, the fully awakened Buddha rolled forth the supreme Wheel of Dhamma. And that wheel cannot be rolled back by any ascetic or brahmin or god or Māra or divinity or by anyone in the world.",
			},
		},
		RepeatPenalty:    1.02,
		RepeatLastN:      512,
		TopK:             0,
		TopP:             1.0,
		MinP:             0.02,
		XtcThreshold:     1.0,
		Temperature:      1.0,
		DryMultiplier:    0.8,
		DryBase:          1.75,
		DryAllowedLength: 2,
		DryPenaltyLastN:  512,
		Stream:           true,
	}

	// Serialize the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("Error serializing request body: %v\n", err)
		os.Exit(1)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", *serverURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		os.Exit(1)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer no-key")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending HTTP request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Print the response
	fmt.Printf("Response status: %s\n", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Response body:\n%s\n", string(body))
}
