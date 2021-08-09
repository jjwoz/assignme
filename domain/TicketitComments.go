package dto

import (
	"database/sql"
	"time"
)

type TicketitComments struct {
	ID        int            `db:"id"`
	UserID    int            `db:"user_id"`
	TicketID  int            `db:"ticket_id"`
	CreatedAt int            `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	Content   sql.NullString `db:"content"`
	Html      sql.NullString `db:"html"`
}
