package main

import (
	"strings"
	"text/template"

	"github.com/irodavlas/api-gateway/handler"

	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Note taking app, for now just make micro service for it.
//can add title to task, description, eta of completion, date of insertion
//will have UI showing title, eta, date

const (
	user_service_url = "http://users-service:5000"
)

func main() {
	e := echo.New()
	e.Static("/assets", "../frontend/assets")
	/*
		err := godotenv.Load("app.env") // Loads variables from .env file
		if err != nil {
			e.Logger.Fatal(err)
		}
	*/
	// needs middleware to check presence of jwt token otw redirect
	secret_key := os.Getenv("SECRET_KEY") //secret key for JWT tokens
	auth_handler := handler.NewAuthService(secret_key)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Path(), "/assets/")
		},
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	protected := e.Group("/auth/*")
	protected.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secret_key),
	}))
	//e.GET("/",)
	e.GET("/", handler.HandleIndexPage)

	e.GET("/events", func(c echo.Context) error {
		templ, _ := template.New("").ParseFiles("../frontend/templates/events.html", "../frontend/templates/layout.html")
		return templ.ExecuteTemplate(c.Response().Writer, "base", nil)

	})
	e.POST("/login", auth_handler.Login)
	e.POST("/sign", auth_handler.SignIn)
	//protected.GET("", router_proxy(auth_service_url))
	e.Logger.Fatal(e.Start(":8000"))
}
