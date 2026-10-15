package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	apiKey     string
	baseUrl    string
	httpClient *http.Client
}

func New(apiKey, baseUrl string) *Client {
	return &Client{
		apiKey:     apiKey,
		baseUrl:    baseUrl,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetWeather(ctx context.Context, city string) (*WeatherResponse, error) {
	//TODO: fix for cities with multiple words
	reqUrl := fmt.Sprintf("%s/weather?q=%s&appid=%s&units=metric", c.baseUrl, city, c.apiKey)
	fmt.Println(reqUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get forecast: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var wr WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&wr); err != nil {
		return nil, err
	}

	return &wr, nil
}
