package dto

type TicketitStatuses struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Color int    `db:"color"`
}
