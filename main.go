package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/1c3fr34k/ticketsystem_backend/database"
	"github.com/1c3fr34k/ticketsystem_backend/middlewares"
	"github.com/1c3fr34k/ticketsystem_backend/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/1c3fr34k/ticketsystem_backend/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
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

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		petstore.swagger.io
// @BasePath	/v2
func main() {

	e := echo.New()

	db, err := database.Connect(dsn)
	if err != nil {
		panic(err)
	}

	e.Use(middlewares.ContextDB(db))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Access-Control-Allow-Origin"},
	}))

	routes.Routes(e.Group("/api"))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("routes.json", data, 0644)

	e.Logger.Fatal(e.Start(":8080"))

}
