package tickets

import "github.com/labstack/echo/v4"

type TicketsRouter struct {
}

func (handler TicketsRouter) Init(g *echo.Group) {
	g.GET("", handler.GetTickets)
	g.POST("/create", handler.CreateTicket)
	g.GET("/:id", handler.GetTicket)
	g.PUT("/:id", handler.UpdateTicket)
	g.PUT("/:id/close", handler.CloseTicket)
	g.PUT("/:id/open", handler.OpenTicket)
}
