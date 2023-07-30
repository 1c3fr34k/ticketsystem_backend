package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/1c3fr34k/ticketsystem_backend/database"
	"github.com/1c3fr34k/ticketsystem_backend/middlewares"
	"github.com/1c3fr34k/ticketsystem_backend/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (
	env_host     string
	env_user     string
	env_password string
	env_database string
	env_port     string
	env_sslmode  string
	dsn          string
)

func init() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		panic(err)
	}

	env_host = os.Getenv("DB_HOST")
	env_user = os.Getenv("DB_USER")
	env_password = os.Getenv("DB_PASSWORD")
	env_database = os.Getenv("DB_DATABASE")
	env_port = os.Getenv("DB_PORT")
	env_sslmode = os.Getenv("DB_SSLMODE")
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", env_host, env_user, env_password, env_database, env_port, env_sslmode)
}

func main() {

	e := echo.New()

	db, err := database.Connect(dsn)
	if err != nil {
		panic(err)
	}

	e.Use(middlewares.ContextDB(db))

	routes.Routes(e.Group("/api"))

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("routes.json", data, 0644)

	e.Logger.Fatal(e.Start(":8080"))

}
