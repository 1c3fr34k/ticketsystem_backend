package tickets

import (
	"net/http"

	"github.com/1c3fr34k/ticketsystem_backend/database/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (TicketsRouter) GetTicket(c echo.Context) error {
	id := c.Param("id")
	var ticket models.Ticket
	db, _ := c.Get("db").(*gorm.DB)
	db.First(&ticket, id)

	return c.JSON(http.StatusOK, ticket)
}

func (TicketsRouter) GetTickets(c echo.Context) error {
	var tickets []models.Ticket
	db, _ := c.Get("db").(*gorm.DB)
	db.Find(&tickets)

	return c.JSON(http.StatusOK, tickets)
}

func (TicketsRouter) CreateTicket(c echo.Context) error {
	var ticket models.Ticket
	c.Bind(&ticket)
	db, _ := c.Get("db").(*gorm.DB)
	db.Create(&ticket)

	return c.JSON(http.StatusOK, ticket)
}

func (TicketsRouter) UpdateTicket(c echo.Context) error {
	id := c.Param("id")
	var ticket models.Ticket
	db, _ := c.Get("db").(*gorm.DB)
	db.First(&ticket, id)
	c.Bind(&ticket)
	db.Save(&ticket)

	return c.JSON(http.StatusOK, ticket)
}

func (TicketsRouter) CloseTicket(c echo.Context) error {
	id := c.Param("id")
	var ticket models.Ticket
	db, _ := c.Get("db").(*gorm.DB)
	db.First(&ticket, id)
	ticket.IsOpen = false
	db.Save(&ticket)
	return c.JSON(http.StatusOK, ticket)
}

func (TicketsRouter) OpenTicket(c echo.Context) error {
	id := c.Param("id")
	var ticket models.Ticket
	db, _ := c.Get("db").(*gorm.DB)
	db.First(&ticket, id)
	ticket.IsOpen = true
	db.Save(&ticket)
	return c.JSON(http.StatusOK, ticket)
}
