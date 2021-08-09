package dto

type TicketitPriorities struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Color int    `db:"color"`
}
