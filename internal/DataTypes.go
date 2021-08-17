package internal

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type DataTypes struct {
	ID                  int            `db:"id"`
	Name                sql.NullString `db:"name"`
	Slug                sql.NullString `db:"slug"`
	DisplayNameSingular sql.NullString `db:"display_name_singular"`
	DisplayNamePlural   sql.NullString `db:"display_name_plural"`
	Icon                sql.NullString `db:"icon"`
	ModelName           sql.NullString `db:"model_name"`
	PolicyName          sql.NullString `db:"policy_name"`
	Controller          sql.NullString `db:"controller"`
	Description         sql.NullString `db:"description"`
	GeneratePermissions sql.NullInt64  `db:"generate_permissions"`
	ServerSIDe          sql.NullInt64  `db:"server_side"`
	Details             sql.NullString `db:"details"`
	CreatedAt           mysql.NullTime `db:"created_at"`
	UpdatedAt           mysql.NullTime `db:"updated_at"`
}
