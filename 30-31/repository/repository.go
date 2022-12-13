package repsitory

import (
	"context"

	"skillbox/30-31/models"
)

// PostRepo explain...
type PostRepo interface {
	Fetch(ctx context.Context, Name string) ([]*models.RequestSelect, error)
	Create(ctx context.Context, p *models.RequestSelect) (int64, error)
	Delete(ctx context.Context, id int64) (bool, error)
	Update(ctx context.Context, p *models.RequestMakeFriend) (*models.RequestMakeFriend, error)
}
