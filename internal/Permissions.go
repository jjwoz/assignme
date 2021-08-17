package internal

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type Permissions struct {
	ID             int            `db:"id"`
	PermissionsKey sql.NullString `db:"permissions_key"`
	TableName      sql.NullString `db:"table_name"`
	CreatedAt      mysql.NullTime `db:"created_at"`
	UpdatedAt      mysql.NullTime `db:"updated_at"`
}
