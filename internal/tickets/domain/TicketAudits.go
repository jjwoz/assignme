package domain

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type TicketAudits struct {
	ID        int            `db:"id"`
	Operation sql.NullString `db:"operation"`
	UserID    sql.NullInt64  `db:"user_id"`
	TicketID  sql.NullInt64  `db:"ticket_id"`
	CreatedAt mysql.NullTime `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
}
