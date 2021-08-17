package internal

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type MenuItems struct {
	ID            int            `db:"id"`
	MenuID        sql.NullInt64  `db:"menu_id"`
	Title         sql.NullString `db:"title"`
	URL           sql.NullString `db:"url"`
	Target        sql.NullString `db:"target"`
	IconClass     sql.NullString `db:"icon_class"`
	Color         sql.NullString `db:"color"`
	ParentID      sql.NullInt64  `db:"parent_id"`
	MenuItemOrder sql.NullInt64  `db:"menu_item_order"`
	CreatedAt     mysql.NullTime `db:"created_at"`
	UpdatedAt     mysql.NullTime `db:"updated_at"`
	Route         sql.NullString `db:"route"`
	Parameters    sql.NullString `db:"parameters"`
}
