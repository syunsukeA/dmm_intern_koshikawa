package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Fetch status which has specified id
	FindByID(ctx context.Context, id int64) (*object.Status, error)
	// Save status
	SaveStatus(ctx context.Context, so *object.Status) (*object.Status, error)
	// TODO: Add Other APIs
}
