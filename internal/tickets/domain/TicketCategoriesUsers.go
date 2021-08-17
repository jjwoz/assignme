package domain

import (
	"database/sql"
)

type TicketCategoriesUsers struct {
	ID         int           `db:"id"`
	CategoryID sql.NullInt64 `db:"category_id"`
	UserID     sql.NullInt64 `db:"user_id"`
}
