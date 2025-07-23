package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xjncx/people-info-api/internal/config"
)

type AgifyClient struct {
	httpClient *http.Client
	URL        string
}

func NewAgifyClient(cfg *config.Config) *AgifyClient {
	return &AgifyClient{
		httpClient: &http.Client{
			Timeout: cfg.HTTPTimeout,
			Transport: &http.Transport{
				MaxIdleConns:    cfg.HTTPMaxIdleConns,
				IdleConnTimeout: cfg.HTTPIdleConnTimeout,
			},
		},
		URL: cfg.AgifyURL,
	}
}

type AgifyResponse struct {
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (c *AgifyClient) GetAge(ctx context.Context, name string) (int, error) {
	url := fmt.Sprintf("%s/?name=%s", c.URL, name)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("genderize API returned status %d", resp.StatusCode)
	}

	var result AgifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Age, nil
}

func (c *AgifyClient) Enrich(ctx context.Context, name string) (*EnrichmentData, error) {
	age, err := c.GetAge(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get age from Agify: %w", err)
	}

	return &EnrichmentData{
		Age: &age,
	}, nil
}

func (c *AgifyClient) Name() string {
	return "Agify"
}
