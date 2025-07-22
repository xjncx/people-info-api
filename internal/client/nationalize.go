package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xjncx/people-info-api/internal/config"
	"github.com/xjncx/people-info-api/internal/model"
)

type NationalizeClient struct {
	httpClient *http.Client
	URL        string
}

type Country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type NationalizeResponse struct {
	Name    string    `json:"name"`
	Country []Country `json:"country"`
}

func NewNationalizeClient(cfg *config.Config) *NationalizeClient {
	return &NationalizeClient{
		httpClient: &http.Client{
			Timeout: cfg.HTTPTimeout,
			Transport: &http.Transport{
				MaxIdleConns:    cfg.HTTPMaxIdleConns,
				IdleConnTimeout: cfg.HTTPIdleConnTimeout,
			},
		},
		URL: cfg.NationalizeURL,
	}
}

func (c *NationalizeClient) GetNationality(ctx context.Context, name string) (string, error) {

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

	var result NationalizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Country) == 0 {
		return "", nil
	}

	maxProb := 0.0
	countryID := ""
	for _, c := range result.Country {
		if c.Probability > maxProb {
			maxProb = c.Probability
			countryID = c.CountryID
		}
	}

	return countryID, nil

}

func (c *NationalizeClient) Enrich(ctx context.Context, name string, person *model.Person) error {
	n, err := c.GetNationality(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to get age from Agify: %w", err)
	}
	person.Nationality = n
	return nil
}

func (c *NationalizeClient) Name() string {
	return "Nationalize"
}
