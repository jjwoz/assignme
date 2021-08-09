package dto

import (
	"time"
)

type TicketitAudits struct {
	ID        int       `db:"id"`
	Operation string    `db:"operation"`
	UserID    int       `db:"user_id"`
	TicketID  int       `db:"ticket_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
