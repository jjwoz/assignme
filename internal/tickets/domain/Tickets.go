package domain

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type NullTime sql.NullTime

type Tickets struct {
	ID          int            `db:"id"`
	Subject     sql.NullString `db:"subject"`
	Content     sql.NullString `db:"content"`
	Html        sql.NullString `db:"html"`
	StatusID    sql.NullInt64  `db:"status_id"`
	PriorityID  sql.NullInt64  `db:"priority_id"`
	AssigneeID  sql.NullInt64  `db:"assignee_id"`
	CreatorID   int            `db:"creator_id"`
	CategoryID  sql.NullInt64  `db:"category_id"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	CompletedAt time.Time      `db:"completed_at"`
}

type TicketResponse struct {
	ID          int      `json:"id"`
	Subject     *string  `json:"subject"`
	Content     *string  `json:"content"`
	Html        *string  `json:"html,omitempty"`
	Status      string   `json:"status"`
	Priority    string   `json:"priority"`
	Assignee    *string  `json:"assignee,omitempty"`
	Creator     string   `json:"creator"`
	Category    string   `json:"category"`
	CreatedAt   NullTime `json:"createdAt,omitempty"`
	CompletedAt NullTime `json:"completedAt,omitempty"`
}

func (nt *NullTime) Scan(value interface{}) error {
	var i sql.NullTime
	if err := i.Scan(value); err != nil {
		return err
	}
	if reflect.TypeOf(value) == nil {
		*nt = NullTime{Time: i.Time, Valid: false}
	} else {
		*nt = NullTime{Time: i.Time, Valid: true}
	}
	return nil
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}
