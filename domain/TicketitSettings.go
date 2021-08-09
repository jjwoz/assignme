package dto

import (
	"database/sql"
	"time"
)

type TicketitSettings struct {
	ID             int            `db:"id"`
	Lang           sql.NullString `db:"lang"`
	Slug           string         `db:"slug"`
	Value          sql.NullString `db:"value"`
	DefaultSetting sql.NullString `db:"default_setting"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}
