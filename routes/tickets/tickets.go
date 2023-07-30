package tickets

import "github.com/labstack/echo/v4"

type TicketsRouter struct {
}

func (ctrl TicketsRouter) Init(g *echo.Group) {
	g.GET("", ctrl.GetTickets)
	g.POST("/create", ctrl.CreateTicket)
	g.GET("/:id", ctrl.GetTicket)
	g.PUT("/:id", ctrl.UpdateTicket)
	g.PUT("/:id/close", ctrl.CloseTicket)
	g.PUT("/:id/open", ctrl.OpenTicket)
}
