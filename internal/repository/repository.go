package repository

import (
	"RaribleAPI/internal/model"
	"context"
)

type NFT interface {
	GetRarity(ctx context.Context, cat *model.Cat) error
	GetOwnershipById(ctx context.Context, cat *model.Cat) error
}
