package internal

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type Translations struct {
	ID         int            `db:"id"`
	TableName  sql.NullString `db:"table_name"`
	ColumnName sql.NullString `db:"column_name"`
	ForeignKey sql.NullInt64  `db:"foreign_key"`
	Locale     sql.NullString `db:"locale"`
	Value      sql.NullString `db:"value"`
	CreatedAt  mysql.NullTime `db:"created_at"`
	UpdatedAt  mysql.NullTime `db:"updated_at"`
}
