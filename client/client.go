package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/romankravchuk/money-converter/types"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) ConvertCurrency(ctx context.Context, from, to string, amount float64) (*types.ConvertResponse, error) {
	req, err := http.NewRequest("GET", c.endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service responsed with non OK status code: %s", httpErr["error"])
	}

	convertResp := new(types.ConvertResponse)
	if err := json.NewDecoder(resp.Body).Decode(convertResp); err != nil {
		return nil, err
	}

	return convertResp, nil
}
