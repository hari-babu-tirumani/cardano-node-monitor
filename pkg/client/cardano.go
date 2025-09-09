package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CardanoClient struct {
	baseURL    string
	httpClient *http.Client
}

type TipResponse struct {
	BlockNo      uint64        `json:"blockNo"`
	SlotNo       uint64        `json:"slotNo"`
	Hash         string        `json:"hash"`
	Epoch        uint64        `json:"epoch"`
	ResponseTime time.Duration `json:"-"`
}

type NetworkResponse struct {
	LocalTip TipResponse `json:"localTip"`
	NetworkTip TipResponse `json:"networkTip"`
	NodeEra  string      `json:"nodeEra"`
	SyncProgress string  `json:"syncProgress"`
}

type NetworkStakeResponse struct {
	Pools map[string]interface{} `json:"pools"`
}

func NewCardanoClient(baseURL string) *CardanoClient {
	return &CardanoClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *CardanoClient) GetTip(ctx context.Context) (*TipResponse, error) {
	start := time.Now()
	
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/api/rest/v0/tip", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tip TipResponse
	if err := json.Unmarshal(body, &tip); err != nil {
		return nil, err
	}

	// Record response time
	duration := time.Since(start)
	tip.ResponseTime = duration

	return &tip, nil
}

func (c *CardanoClient) GetNetwork(ctx context.Context) (*NetworkResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/api/rest/v0/network", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("network API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var network NetworkResponse
	if err := json.Unmarshal(body, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

func (c *CardanoClient) IsHealthy(ctx context.Context) bool {
	_, err := c.GetTip(ctx)
	return err == nil
}