package dto

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type Ticketit struct {
	ID          int            `db:"id"`
	Subject     string         `db:"subject"`
	Content     *string        `db:"content"`
	Html        *string        `db:"html" json:",omitempty"`
	StatusID    int            `db:"status_id"`
	PriorityID  int            `db:"priority_id"`
	UserID      int            `db:"user_id"`
	AgentID     int            `db:"agent_id"`
	CategoryID  int            `db:"category_id"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   int            `db:"updated_at"`
	CompletedAt mysql.NullTime `db:"completed_at"`
}
