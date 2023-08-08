package users

import "github.com/labstack/echo/v4"

type UsersRouter struct {
}

func (handler UsersRouter) Init(g *echo.Group) {
	g.POST("/create", handler.CreateUser)
	g.GET("/:id", handler.GetUser)
	g.GET("", handler.GetUsers)
}
