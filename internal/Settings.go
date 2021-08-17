package internal

import (
	"database/sql"
)

type Settings struct {
	ID                int            `db:"id"`
	SettingsKey       sql.NullString `db:"settings_key"`
	DisplayName       sql.NullString `db:"display_name"`
	Value             sql.NullString `db:"value"`
	Details           sql.NullString `db:"details"`
	Type              sql.NullString `db:"type"`
	SettingsItemOrder sql.NullInt64  `db:"settings_item_order"`
	SettingsGroup     sql.NullInt64  `db:"settings_group"`
	UserID            sql.NullInt64  `db:"user_id"`
}
