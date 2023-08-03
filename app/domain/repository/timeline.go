package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	// Fetch status which has specified id
	FindByID(ctx context.Context, only_media bool, max_id int64, since_id int64, limit int64) (*object.Timeline, error)
}
