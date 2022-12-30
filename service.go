package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/romankravchuk/money-converter/types"
)

type CurrencyConverter interface {
	Convert(context.Context, string, string, float64) (float64, error)
}

type currencyConverter struct {
	APIUrl string
	APIKey string
}

func NewCurrencyConverter(url, key string) CurrencyConverter {
	return &currencyConverter{
		APIUrl: url,
		APIKey: key,
	}
}

func (s *currencyConverter) Convert(ctx context.Context, from, to string, amount float64) (float64, error) {
	endpoint := fmt.Sprintf("%s?to=%s&from=%s&amount=%.2f", s.APIUrl, to, from, amount)

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("apikey", s.APIKey)
	if err != nil {
		return -1, err
	}

	resp, err := http.DefaultClient.Do(req)
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return -1, err
	}

	var convertResp types.ConvertResponse
	if err := json.NewDecoder(resp.Body).Decode(&convertResp); err != nil {
		return -1, err
	}
	result := convertResp.Result * amount
	return result, nil
}
