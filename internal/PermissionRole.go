package internal

import (
	"database/sql"
)

type PermissionRole struct {
	PermissionID int           `db:"permission_id"`
	RoleID       sql.NullInt64 `db:"role_id"`
}
