package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/margostino/babel-cli/internal/common"
	"github.com/margostino/babel-cli/internal/prompt"
)

var BASE_URL = "https://api.openai.com/v1"
var CHAT_COMPLETION_PATH = "/chat/completions"
var MODEL = "gpt-4o"
var PROMPT_FILE_PATH = "./metadata_enricher.yml"

type LLMTask struct {
	RelativeFilePath string
	Content          []byte
	ApiKey           string
}

type LLMTaskResult struct {
	RelativeFilePath string
	Content          string
	Error            error
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseFormatType string

const (
	JSON_OBJECT ResponseFormatType = "json_object"
	JSON_SCHEMA ResponseFormatType = "json_schema"
)

type Property struct {
	Type string `json:"type"`
}

type JsonSchemaType struct {
	Type                 string              `json:"type"`
	Properties           map[string]Property `json:"properties"`
	Required             []string            `json:"required"`
	AdditionalProperties bool                `json:"additionalProperties"`
}

type JsonSchema struct {
	Name   string         `json:"name"`
	Schema JsonSchemaType `json:"schema"`
}

type ResponseFormat struct {
	Type       ResponseFormatType `json:"type"`
	JsonSchema JsonSchema         `json:"json_schema"`
	Strict     bool               `json:"strict"`
}

type RequestBody struct {
	Model          string         `json:"model"`
	Messages       []Message      `json:"messages"`
	ResponseFormat ResponseFormat `json:"response_format"`
}

type Choice struct {
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
}

type ApiResponse struct {
	Choices []Choice `json:"choices"`
}

func GetRequest(path string, input string) map[string]interface{} {
	systemPrompt, err := prompt.GetMetadataEnricher()
	common.Check(err, "Failed to get system prompt")
	// messages := []Message{
	// 	{
	// 		Role:    "system",
	// 		Content: systemPrompt,
	// 	},
	// 	{
	// 		Role:    "user",
	// 		Content: fmt.Sprintf("File Path: %s", path),
	// 	},
	// 	{
	// 		Role:    "user",
	// 		Content: fmt.Sprintf("File content: %s", input),
	// 	},
	// }
	// return &RequestBody{
	// 	Model:    MODEL,
	// 	Messages: messages,
	// 	ResponseFormat: ResponseFormat{
	// 		Type:   JSON_SCHEMA,
	// 		Strict: true,
	// 		JsonSchema: JsonSchema{
	// 			Name: "metadata",
	// 			Schema: JsonSchemaType{
	// 				Type: "object",
	// 				Properties: map[string]Property{
	// 					"category": {
	// 						Type: "string",
	// 					},
	// 					"highlights": {
	// 						Type: "array",
	// 						Items: JsonSchemaType{
	// 							Type:                 "string",
	// 							AdditionalProperties: false,
	// 						},
	// 					},
	// 					"keywords": {
	// 						Type: "array",
	// 					},
	// 					"path": {
	// 						Type: "string",
	// 					},
	// 					"references": {
	// 						Type: "array",
	// 					},
	// 					"related_links": {
	// 						Type: "array",
	// 					},
	// 					"summary": {
	// 						Type: "string",
	// 					},
	// 					"tags": {
	// 						Type: "array",
	// 					},
	// 				},
	// 				Required: []string{
	// 					"category",
	// 					"highlights",
	// 					"keywords",
	// 					"path",
	// 					"references",
	// 					"related_links",
	// 					"summary",
	// 					"tags",
	// 				},
	// 				AdditionalProperties: false,
	// 			},
	// 		},
	// 	},
	// }
	return map[string]interface{}{
		"model": MODEL,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": systemPrompt,
			},
			{
				"role":    "user",
				"content": fmt.Sprintf("File Path: %s", path),
			},
			{
				"role":    "user",
				"content": fmt.Sprintf("File content: %s", input),
			},
		},
		"response_format": map[string]interface{}{
			"type": "json_schema",
			"json_schema": map[string]interface{}{
				"name":   "metadata",
				"strict": true,
				"schema": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"category":      map[string]string{"type": "string"},
						"highlights":    map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string", "additionalProperties": false}},
						"keywords":      map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string", "additionalProperties": false}},
						"path":          map[string]string{"type": "string"},
						"references":    map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string", "additionalProperties": false}},
						"related_links": map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string", "additionalProperties": false}},
						"summary":       map[string]string{"type": "string"},
						"tags":          map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string", "additionalProperties": false}},
					},
					"required": []string{
						"category", "highlights", "keywords", "path",
						"references", "related_links", "summary", "tags",
					},
					"additionalProperties": false,
				},
			},
		},
	}
}

func GetChatCompletionForMetadata(apiKey string, relativeFilePath string, content string) (string, error) {
	apiURL := BASE_URL + CHAT_COMPLETION_PATH

	requestBody := GetRequest(relativeFilePath, content)

	jsonData, err := json.Marshal(requestBody)
	// fmt.Println("Request body:", string(jsonData))
	common.Check(err, "Failed to marshal OpenAI request body")

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, apiURL, bytes.NewBuffer(jsonData))
	common.Check(err, "Failed to create OpenAI request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	time.Sleep(3 * time.Second)
	resp, err := client.Do(req)
	common.Check(err, "Failed to send OpenAI request")
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return "", fmt.Errorf("OpenAI request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	common.Check(err, "Failed to read OpenAI response body")

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	common.Check(err, "Failed to unmarshal OpenAI response body")

	if len(apiResponse.Choices) == 0 {
		return "", fmt.Errorf("No choices found in OpenAI response")
	}

	return apiResponse.Choices[0].Message.Content, nil
}

func ProcessChatCompletionForMetadataTask(task LLMTask) (string, error) {
	apiURL := BASE_URL + CHAT_COMPLETION_PATH

	requestBody := GetRequest(task.RelativeFilePath, string(task.Content))

	jsonData, err := json.Marshal(requestBody)
	common.Check(err, "Failed to marshal OpenAI request body")

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, apiURL, bytes.NewBuffer(jsonData))
	common.Check(err, "Failed to create OpenAI request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+task.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	common.Check(err, "Failed to send OpenAI request")
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return "", fmt.Errorf("OpenAI request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	common.Check(err, "Failed to read OpenAI response body")

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	common.Check(err, "Failed to unmarshal OpenAI response body")

	if len(apiResponse.Choices) == 0 {
		return "", fmt.Errorf("No choices found in OpenAI response")
	}

	return apiResponse.Choices[0].Message.Content, nil
}

func ChatCompletionForMetadataWorker(taskChan <-chan LLMTask, taskResultChan chan<- LLMTaskResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		content, err := ProcessChatCompletionForMetadataTask(task)
		taskResultChan <- LLMTaskResult{RelativeFilePath: task.RelativeFilePath, Content: content, Error: err}
	}
}
