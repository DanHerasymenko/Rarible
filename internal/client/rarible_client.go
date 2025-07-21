package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type RaribleClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func NewRaribleClient(baseURL, apiKey string) *RaribleClient {
	return &RaribleClient{
		baseURL: baseURL,
		apiKey:  apiKey,
		client:  &http.Client{},
	}
}

// GetNFTOwnerships retrieves a list of NFT ownerships by id
func (rc *RaribleClient) GetNFTOwnerships(ctx context.Context, ownershipId string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("%s/ownerships/%s", rc.baseURL, url.PathEscape(ownershipId))
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	if rc.apiKey != "" {
		req.Header.Set("X-API-KEY", rc.apiKey)
	}
	resp, err := rc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTraitRarities retrieves NFT trait rarities
func (rc *RaribleClient) GetTraitRarities(ctx context.Context, collection string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("%s/collections/%s/traits/rarity", rc.baseURL, url.PathEscape(collection))
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	if rc.apiKey != "" {
		req.Header.Set("X-API-KEY", rc.apiKey)
	}
	resp, err := rc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
