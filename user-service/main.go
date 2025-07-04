package main

import (
	"github.com/irodavlas/user-service/database"
	"github.com/irodavlas/user-service/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// this microservice will handle private routes for user creation, update and removal
func main() {

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))
	e.Use(middleware.Recover())
	mongo_URI := "mongodb://172.17.0.1:27018" //os.Getenv("MONGO_URI")
	println("mongo url:", mongo_URI)
	db, err := database.NewDatabaseConnection(mongo_URI, "usersdb", "users")
	if err != nil {
		e.Logger.Fatal(err)
	}
	err = db.PingDatabase()
	if err != nil {
		e.Logger.Fatal(err)
	}
	service := handler.NewUserService(db)
	//PUT to update the resources, POST to create
	e.POST("/create", service.CreateUser)
	e.POST("/delete", service.DeleteUser)
	e.GET("/get", service.GetUser)
	// must be updated later on
	e.PUT("/update", service.UpdateUser)

	e.Logger.Fatal(e.Start(":5000"))
}
