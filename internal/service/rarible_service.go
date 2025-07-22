package service

import (
	"RaribleAPI/internal/client"
	"RaribleAPI/internal/model"
	"context"
)

type RaribleService struct {
	client *client.RaribleClient
}

func NewRaribleService(client *client.RaribleClient) *RaribleService {
	return &RaribleService{client: client}
}

func (s *RaribleService) GetNFTOwnerships(ctx context.Context, ownershipId string) (*model.OwnershipResponse, error) {
	return s.client.GetNFTOwnerships(ctx, ownershipId)
}

func (s *RaribleService) GetTraitRarities(ctx context.Context, body model.RarityRequest) (*model.RarityResponse, error) {
	return s.client.GetTraitRaritiesPOST(ctx, body)
}
