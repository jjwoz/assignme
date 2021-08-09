package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	dto "github.com/jjwoz/assignme/domain"
	"log"
)

const (
	insertTicket     = "INSERT INTO ticketit(subject, content, html, status_id, priority_id, user_id, agent_id, category_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
	selectAllTickets = "select id, subject, content, html, status_id, priority_id, user_id, agent_id, category_id from ticketit order by id;"
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
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", localConfig.User, localConfig.Pass, localConfig.DatabaseName))
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("ticket", func(c *fiber.Ctx) error {
		//new ticket object
		t := new(dto.Ticketit)
		// Parse body into struct
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		//Insert into DB
		_, err := db.Query(insertTicket, t.Subject, t.Content, t.Html, t.StatusID, t.PriorityID, t.UserID, t.AgentID, t.CategoryID)
		if err != nil {
			return err
		}
		//return
		return c.JSON(t)

	})

	app.Get("tickets", func(c *fiber.Ctx) error {
		var tickets []dto.Ticketit
		rows, err := db.Query(selectAllTickets)
		defer rows.Close()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		//select id, subject, content, html, status_id, priority_id, user_id, agent_id, category_id from ticketit order by id;
		for rows.Next() {
			t := dto.Ticketit{}
			if err := rows.Scan(&t.ID, &t.Subject, &t.Content, &t.Html, &t.StatusID, &t.PriorityID, &t.UserID, &t.AgentID, &t.CategoryID); err != nil {
				return err
			}

			tickets = append(tickets, t)
		}
		return c.JSON(tickets)
	})

	log.Fatal(app.Listen(":3000"))
}
