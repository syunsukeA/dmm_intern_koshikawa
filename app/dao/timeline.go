package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_"log"
	_ "time"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	timeline struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

// FindByID : ID関連から複数投稿を取得しtimelineに整形
// ToDo: only_mediaのtrue/falseで処理を分岐させる
func (r *timeline) FindByID(ctx context.Context, only_media bool, max_id int64, since_id int64, limit int64) (*object.Timeline, error) {
	timeline := new(object.Timeline)
	/*
		QueryxContextは該当行がない場合もerrがnilで帰ってくる
		rowsのfor文に入ったかどうかをforFlagで判定する
	*/
	rows_is_empty := true

	// DBからデータ取得
	rows, err := r.db.QueryxContext(
		ctx,
		`select
			*
		from
			status
		where 
			id<=? and id>=?
		limit ?`,
		max_id, since_id, limit,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find id from db: %w", err)
	}
	// 各投稿ごとにAccountと紐付けてTimelineElementにまとめる
	for rows.Next() {
		rows_is_empty = false
		// ToDo: N+1の解消
		tl_element := new(object.TimelineElement)
		os := new(object.Status)
		oa := new(object.Account)
		err = rows.StructScan(os)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("failed to find id from db: %w", err)
		}
		err = r.db.QueryRowxContext(ctx, "select * from account where id = ?", os.AccountID).StructScan(oa)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("failed to find id from db: %w", err)
		}
		tl_element.Status = os
		tl_element.Account = oa
		timeline.Timeline = append(timeline.Timeline, tl_element)
	}
	if rows_is_empty {
		return nil, nil
	}

	return timeline, nil
}

// AuthFindByID : ID関連から複数投稿を取得しtimelineに整形
// ToDo: only_mediaのtrue/falseで処理を分岐させる
func (r *timeline) AuthFindByID(ctx context.Context, ao *object.Account, only_media bool, max_id int64, since_id int64, limit int64) (*object.Timeline, error) {
	timeline := new(object.Timeline)
	/*
		QueryxContextは該当行がない場合もerrがnilで帰ってくる
		rowsのfor文に入ったかどうかをforFlagで判定する
	*/
	rows_is_empty := true

	// DBからデータ取得
	rows, err := r.db.QueryxContext(
		ctx,
		`select
			*
		from
			status
		where 
			id<=? and id>=? and account_id=?
		limit ?`,
		max_id, since_id, ao.ID, limit,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find id from db: %w", err)
	}
	// 各投稿ごとにAccountと紐付けてTimelineElementにまとめる
	for rows.Next() {
		rows_is_empty = false
		tl_element := new(object.TimelineElement)
		so := new(object.Status)
		err = rows.StructScan(so)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("failed to find id from db: %w", err)
		}
		tl_element.Status = so
		tl_element.Account = ao
		timeline.Timeline = append(timeline.Timeline, tl_element)
	}
	
	if rows_is_empty {
		return nil, nil
	}
	return timeline, nil
}
