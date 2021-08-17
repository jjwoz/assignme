package domain

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type TicketComments struct {
	ID        int            `db:"id"`
	UserID    sql.NullInt64  `db:"user_id"`
	TicketID  sql.NullInt64  `db:"ticket_id"`
	CreatedAt mysql.NullTime `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
	Content   sql.NullString `db:"content"`
	Html      sql.NullString `db:"html"`
}
