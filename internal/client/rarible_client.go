package client

import (
	"RaribleAPI/internal/logger"
	"RaribleAPI/internal/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func (rc *RaribleClient) GetNFTOwnerships(ctx context.Context, ownershipId string) (*model.OwnershipResponse, error) {

	endpoint := fmt.Sprintf("%s/ownerships/%s", rc.baseURL, url.PathEscape(ownershipId))
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	if rc.apiKey != "" {
		req.Header.Set("X-API-KEY", rc.apiKey)
	}
	resp, err := rc.client.Do(req)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		logger.Error(ctx, fmt.Errorf("unexpected status: %s", resp.Status))
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var result model.OwnershipResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	return &result, nil
}

func (rc *RaribleClient) GetTraitRaritiesPOST(ctx context.Context, body model.RarityRequest) (*model.RarityResponse, error) {

	endpoint := fmt.Sprintf("%s/items/traits/rarity", rc.baseURL)
	b, err := json.Marshal(body)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(b))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	if rc.apiKey != "" {
		req.Header.Set("X-API-KEY", rc.apiKey)
	}

	resp, err := rc.client.Do(req)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		logger.Error(ctx, fmt.Errorf("unexpected status: %s", resp.Status))
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var result model.RarityResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	return &result, nil
}
