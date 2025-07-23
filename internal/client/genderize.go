package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xjncx/people-info-api/internal/config"
)

type GenderizeClient struct {
	httpClient *http.Client
	URL        string
}

func NewGenderizeClient(cfg *config.Config) *GenderizeClient {
	return &GenderizeClient{
		httpClient: &http.Client{
			Timeout: cfg.HTTPTimeout,
			Transport: &http.Transport{
				MaxIdleConns:    cfg.HTTPMaxIdleConns,
				IdleConnTimeout: cfg.HTTPIdleConnTimeout,
			},
		},
		URL: cfg.GenderizeURL,
	}
}

type GenderizeResponse struct {
	Gender      string  `json:"gender"`
	Name        string  `json:"name"`
	Probability float64 `json:"probability"`
}

func (c *GenderizeClient) GetGender(ctx context.Context, name string) (string, error) {
	url := fmt.Sprintf("%s/?name=%s", c.URL, name)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("genderize API returned status %d", resp.StatusCode)
	}

	var result GenderizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Gender, nil
}

func (c *GenderizeClient) Enrich(ctx context.Context, name string) (*EnrichmentData, error) {
	gender, err := c.GetGender(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get gender from Genderize: %w", err)
	}

	return &EnrichmentData{
		Gender: &gender,
	}, nil
}

func (c *GenderizeClient) Name() string {
	return "Genderize"
}
