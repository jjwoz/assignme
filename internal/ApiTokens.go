package internal

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type ApiTokens struct {
	ID         int            `db:"id"`
	UserID     int            `db:"user_id"`
	Name       string         `db:"name"`
	Token      int            `db:"token"`
	Metadata   string         `db:"metadata"`
	Transient  int            `db:"transient"`
	LastUsedAt mysql.NullTime `db:"last_used_at"`
	ExpiresAt  mysql.NullTime `db:"expires_at"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
}
