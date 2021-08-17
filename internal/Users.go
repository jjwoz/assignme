package internal

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Users struct {
	ID                      int            `db:"id"`
	Name                    string         `db:"name"`
	Email                   string         `db:"email"`
	EmailVerifiedAt         mysql.NullTime `db:"email_verified_at"`
	Password                string         `db:"password"`
	RememberToken           sql.NullString `db:"remember_token"`
	PhotoURL                sql.NullString `db:"photo_url"`
	UsesTwoFactorAuth       int            `db:"uses_two_factor_auth"`
	AuthyID                 sql.NullString `db:"authy_id"`
	CountryCode             sql.NullString `db:"country_code"`
	Phone                   sql.NullString `db:"phone"`
	StripeID                sql.NullString `db:"stripe_id"`
	TwoFactorResetCode      sql.NullString `db:"two_factor_reset_code"`
	CurrentTeamID           sql.NullInt64  `db:"current_team_id"`
	CurrentBillingPlan      sql.NullString `db:"current_billing_plan"`
	CardBrand               sql.NullString `db:"card_brand"`
	CardLastFour            sql.NullString `db:"card_last_four"`
	CardCountry             sql.NullString `db:"card_country"`
	BillingAddress          sql.NullString `db:"billing_address"`
	BillingAddressLine2     sql.NullString `db:"billing_address_line_2"`
	BillingCity             sql.NullString `db:"billing_city"`
	BillingState            sql.NullString `db:"billing_state"`
	BillingZip              sql.NullString `db:"billing_zip"`
	BillingCountry          sql.NullString `db:"billing_country"`
	VatID                   sql.NullString `db:"vat_id"`
	ExtraBillingInformation sql.NullString `db:"extra_billing_information"`
	TrialEndsAt             mysql.NullTime `db:"trial_ends_at"`
	LastReadAnnouncementsAt mysql.NullTime `db:"last_read_announcements_at"`
	CreatedAt               time.Time      `db:"created_at"`
	UpdatedAt               time.Time      `db:"updated_at"`
}
