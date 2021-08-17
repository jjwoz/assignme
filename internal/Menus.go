package internal

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type Menus struct {
	ID        int            `db:"id"`
	Name      sql.NullString `db:"name"`
	CreatedAt mysql.NullTime `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
}
