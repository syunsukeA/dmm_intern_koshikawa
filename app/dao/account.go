package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_"time"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

func (r *account) SaveAccount(ctx context.Context, ao *object.Account) (*object.Account, error) {
	// obj_accountの情報を基にDBに追加
	err := r.db.QueryRowxContext(ctx, "insert into account(username, password_hash) values(?, ?)", ao.Username, ao.PasswordHash).Err()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to add account to db: %w", err)
	}
	// obj_accountを返す
	return ao, nil
}
