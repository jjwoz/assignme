package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	dto "github.com/jjwoz/assignme/internal/tickets/domain"
	"log"
)

const (
	insertTicket     = "INSERT INTO tickets(subject, content, html, status_id, priority_id, user_id, agent_id, category_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
	selectAllTickets = `select t.id,
       t.subject,
       t.content,
       u1.photo_url as creator_photo_url,
       u2.photo_url as assignee_photo_url,
       t.html,
       ts.name,
       tp.name,
       u2.name      as asignee,
       u1.name      as creator,
       tcc.name,
       t.created_at,
       t.completed_at
from tickets t
         join ticket_statuses ts on t.status_id = ts.id
         join ticket_categories tcc on tcc.id = t.category_id
         join ticket_priorities tp on tp.id = t.priority_id
         join users u1 on t.creator_id = u1.id
         left join users u2 on t.assignee_id = u2.id;`
	updateTicket = "UPDATE tickets SET  subject=?, content=?, html=?, status_id=?, priority_id=?, assignee_id=?, category_id=?, updated_at=NOW(), completed_at=? WHERE id = ?;"
)

var db *sql.DB

// Config represents what will be read in from environment
type Config struct {
	Host         string
	Port         int
	User         string
	Pass         string
	DatabaseName string
}

var localConfig = &Config{
	Host:         "localhost",
	Port:         3306,
	User:         "test",
	Pass:         "password123",
	DatabaseName: "assign",
}

func connect() error {
	var err error
	// Use DSN string to open
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?parseTime=true", localConfig.User, localConfig.Pass, localConfig.DatabaseName))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func main() {
	// Connect with database
	if err := connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/ticket", func(c *fiber.Ctx) error {
		//new ticket object
		t := new(dto.Tickets)
		// Parse body into struct
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		//Insert into DB
		_, err := db.Query(insertTicket, t.Subject, t.Content, t.Html, t.StatusID, t.PriorityID, t.AssigneeID, t.CreatorID, t.CategoryID)
		if err != nil {
			return err
		}
		//return
		return c.JSON(t)

	})

	app.Get("/tickets", func(c *fiber.Ctx) error {
		var tickets []dto.TicketResponse
		rows, err := db.Query(selectAllTickets)
		defer rows.Close()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		//select t.id, t.subject, t.content, t.html, ts.name, tp.name, u.name as asignee, u.name as creator,tcc.name, t.created_at
		//from tickets t join ticket_statuses ts on t.status_id = ts.id
		//join ticket_categories tcc on tcc.id = t.category_id
		//join ticket_priorities tp on tp.id = t.priority_id
		//join users u on t.creator_id = u.id;
		for rows.Next() {
			t := dto.TicketResponse{}
			if err := rows.Scan(&t.ID, &t.Subject, &t.Content, &t.CreatorPhotoUrl, &t.AssigneePhotoUrl, &t.Html, &t.Status, &t.Priority, &t.Assignee, &t.Creator, &t.Category, &t.CreatedAt, &t.CompletedAt); err != nil {
				return err
			}

			tickets = append(tickets, t)
		}
		return c.JSON(tickets)
	})
	// Update record into MySQL
	app.Put("/ticket", func(c *fiber.Ctx) error {
		// New Employee struct
		t := new(dto.Tickets)

		// Parse body into struct
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		// UPDATE tickets SET  subject=?, content=?, html=?, status_id=?, priority_id=?, assignee_id=?, category_id=?, updated_at=NOW(), completed_at=? WHERE id = ?;
		// Update Ticket record in database
		_, err := db.Query(updateTicket, t.Subject, t.Content, t.Html, t.StatusID, t.PriorityID, t.AssigneeID, t.CategoryID, t.CompletedAt, t.ID)
		if err != nil {
			return err
		}

		// Return Employee in JSON format
		return c.Status(201).JSON(t)
	})

	// Delete record from MySQL
	app.Delete("/ticket", func(c *fiber.Ctx) error {
		// New Employee struct
		t := new(dto.Tickets)

		// Parse body into struct
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Delete Employee from database
		_, err := db.Query("DELETE FROM Tickets WHERE id = ?", t.ID)
		if err != nil {
			return err
		}

		// Return Employee in JSON format
		return c.JSON("Deleted")
	})
	log.Fatal(app.Listen(":3004"))
}
