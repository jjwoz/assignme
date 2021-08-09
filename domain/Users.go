package dto

type Users struct {
	ID            int `db:"id"`
	TicketitAdmin int `db:"ticketit_admin"`
	TicketitAgent int `db:"ticketit_agent"`
}
