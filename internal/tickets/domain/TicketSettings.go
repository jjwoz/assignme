package domain

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type TicketSettings struct {
	ID            int            `db:"id"`
	Lang          sql.NullString `db:"lang"`
	Slug          sql.NullString `db:"slug"`
	Value         sql.NullString `db:"value"`
	TicketDefault sql.NullString `db:"ticket_default"`
	CreatedAt     mysql.NullTime `db:"created_at"`
	UpdatedAt     mysql.NullTime `db:"updated_at"`
}
