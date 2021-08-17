package domain

import (
	"database/sql"
)

type TicketRoles struct {
	ID               int           `db:"id"`
	TicketCreatorID  sql.NullInt64 `db:"ticket_creator_id"`
	TicketAssigneeID sql.NullInt64 `db:"ticket_assignee_id"`
}
