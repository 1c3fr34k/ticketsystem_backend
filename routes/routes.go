package routes

import (
	"github.com/1c3fr34k/ticketsystem_backend/routes/tickets"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	tickets.TicketsRouter{}.Init(g.Group("/tickets"))
}
