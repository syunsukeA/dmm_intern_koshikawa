package object

import (
	_"fmt"
	"time"

	_"golang.org/x/crypto/bcrypt"
)

type Status struct {
	// The internal ID of the status
	ID int64 `json:"id,omitempty"`

	// The internal ID of the account
	AccountID int64 `json:"account_id,omitempty"`

	// The content of the status
	Content string `json:"content,omitempty"`

	// URL to the ???
	URL *string `json:"url,omitempty"`

	// The time the status was created
	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`
}
