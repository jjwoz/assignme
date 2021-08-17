package domain

import (
	"database/sql"
)

type TicketPriorities struct {
	ID    int            `db:"id"`
	Name  sql.NullString `db:"name"`
	Color sql.NullInt64  `db:"color"`
}
