package ollama

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ollama/ollama/api"

	"name-for-me/box"
)

type Client struct {
	client *api.Client
}

func NewClient(address string) (*Client, error) {
	base, err := url.Parse(address)
	if err != nil {
		return nil, fmt.Errorf("failed to parse address: %w", err)
	}
	httpClient := &http.Client{
		Timeout: 30 * time.Minute,
	}

	client := api.NewClient(base, httpClient)

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetModels(ctx context.Context) ([]string, error) {
	resp, err := c.client.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %w", err)
	}

	result := make([]string, len(resp.Models))
	for i, model := range resp.Models {
		result[i] = model.Model
	}

	return result, nil
}

func (c *Client) Generate(ctx context.Context, model string, input string, callback func(string, bool)) error {
	if err := c.client.Generate(ctx, &api.GenerateRequest{
		Model:  model,
		Prompt: input,
		Stream: box.New(true),
	}, func(response api.GenerateResponse) error {
		callback(response.Response, response.Done)
		return nil
	}); err != nil {
		return fmt.Errorf("failed to chat: %w", err)
	}

	return nil
}
