package gemini

import "net/http"

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
