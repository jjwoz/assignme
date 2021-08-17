package internal

import (
	"github.com/go-sql-driver/mysql"
)

type GooseDbVersion struct {
	ID        int            `db:"id"`
	VersionID int            `db:"version_id"`
	IsApplied int            `db:"is_applied"`
	Tstamp    mysql.NullTime `db:"tstamp"`
}
