package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_"log"
	_"time"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// FindByID : IDから投稿を取得
func (r *status) FindByID(ctx context.Context, id int64) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find id from db: %w", err)
	}

	return entity, nil
}

func (r *status) SaveStatus(ctx context.Context, so *object.Status) (*object.Status, error) {
	// obj_statusの情報を基にDBに追加
	err := r.db.QueryRowxContext(ctx, "insert into status(account_id, content) values(?, ?)", so.AccountID, so.Content).Err()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to add status to db: %w", err)
	}
	// obj_accountを返す
	return so, nil
}
