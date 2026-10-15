package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	apiKey     string
	baseUrl    string
	model      string
	httpClient *http.Client
}

func New(apiKey, baseUrl, model string) *Client {
	return &Client{
		apiKey:     apiKey,
		baseUrl:    baseUrl,
		model:      model,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetAdviceForWeather(ctx context.Context, city string, desc string, temp, wind, rain float64, clouds uint32) (string, error) {
	reqURL := fmt.Sprintf("%s/chat/completions", c.baseUrl)

	reqBody := ChatCompletionRequest{
		Model: c.model,
		Messages: []Message{
			{
				Role: "user",
				Content: fmt.Sprintf(
					"Give me advice about weather in %s. Temperature: %f, conditions: %s, wind: %f, clouds: %d, rain: %f. Maximum 2 sentences. Advice what to wear and Be friendly and say in Russian.",
					city,
					temp,
					desc,
					wind,
					clouds,
					rain,
				),
			},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		reqURL,
		bytes.NewReader(jsonBody),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode == http.StatusServiceUnavailable {
			return "", fmt.Errorf("service is currently overloaded. Try again later")
		}

		return "", fmt.Errorf("GROQ API error: %s", string(body))
	}

	var rc ChatCompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&rc); err != nil {
		return "", err
	}

	return rc.Choices[0].Message.Content, nil
}
