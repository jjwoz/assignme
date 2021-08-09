package dto

type TicketitCategories struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Color int    `db:"color"`
}
