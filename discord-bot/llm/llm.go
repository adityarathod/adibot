package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"log"
)

type CompletionRequest struct {
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature"`
}

type CompletionResponse struct {
	Choices []struct {
		Text         string `json:"text"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

type LLMClient struct {
	Endpoint string
}

func NewLLMClient(endpoint string) LLMClient {
	return LLMClient{
		Endpoint: endpoint,
	}
}

func (cl *LLMClient) CallModelAPI(prompt string) string {
	req := CompletionRequest{
		Prompt:      fmt.Sprintf("%s\n", prompt),
		Temperature: 0.5,
	}

	// send json request to the llm
	reqBytes, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(cl.Endpoint, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// deserialize the respons
	var completionResp CompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&completionResp)
	if err != nil {
		panic(err)
	}
	if len(completionResp.Choices) == 0 {
		log.Printf("No completion found for the prompt %s", prompt)
		return ""
	}
	response := completionResp.Choices[0].Text
	log.Printf("Prompt: \"%s\", Response: \"%s\"", prompt, response)
	return response
}
