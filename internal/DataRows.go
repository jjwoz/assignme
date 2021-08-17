package internal

import (
	"database/sql"
)

type DataRows struct {
	ID          int            `db:"id"`
	DataTypeID  sql.NullInt64  `db:"data_type_id"`
	Field       sql.NullString `db:"field"`
	Type        sql.NullString `db:"type"`
	DisplayName sql.NullString `db:"display_name"`
	Required    sql.NullInt64  `db:"required"`
	Browse      sql.NullInt64  `db:"browse"`
	ReadData    sql.NullInt64  `db:"read_data"`
	Edit        sql.NullInt64  `db:"edit"`
	AddData     sql.NullInt64  `db:"add_data"`
	DeleteData  sql.NullInt64  `db:"delete_data"`
	Details     sql.NullString `db:"details"`
	OrderData   sql.NullInt64  `db:"order_data"`
}
