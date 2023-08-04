package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	// Save account
	SaveAccount(ctx context.Context, obj_account *object.Account) (*object.Account, error)
	// TODO: Add Other APIs
	AddFollow(ctx context.Context, followee_ao *object.Account, follower_ao *object.Account) (*object.Account, error)
}
