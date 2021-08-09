package dto

type TicketitCategoriesUsers struct {
	ID         int `db:"id"`
	CategoryID int `db:"category_id"`
	UserID     int `db:"user_id"`
}
