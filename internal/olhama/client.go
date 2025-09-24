package olhama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// API interna que faz requisições à API do olhama

type generateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type generateResponse struct {
	Response        string `json:"response"`
	PromptEvalCount int    `json:"prompt_eval_count"`
	EvalCount       int    `json:"eval_count"`
}

type Client struct {
	ollhmaURL  string
	httpClient *http.Client
}

func NewClient(url string) *Client {
	return &Client{
		ollhmaURL:  url,
		httpClient: &http.Client{Timeout: 1 * time.Minute},
	}
}

func (c *Client) Generate(ctx context.Context, prompt string) (*generateResponse, error) {
	// os parametros a seguir devem ser substituidas por variaveis de ambiente carregadas em "config"
	reqPayload := generateRequest{
		Model:  "llama3",
		Prompt: prompt,
		Stream: false,
	}
	payloadBytes, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, fmt.Errorf("error to marshal request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.ollhmaURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("error to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error to contact olhama API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro olhama API return status %s", resp.Status)
	}

	var olhamaResponse generateResponse
	if err := json.NewDecoder(resp.Body).Decode(olhamaResponse); err != nil {
		return nil, fmt.Errorf("error to decode olhama json: %v", err)
	}

	return &olhamaResponse, nil
}
