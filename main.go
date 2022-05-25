package main

import (
	"fmt"

	"github.com/Sleeplessss/first_go_fiber/controllers"
	"github.com/Sleeplessss/first_go_fiber/models"
	"github.com/Sleeplessss/first_go_fiber/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App)  {
	app.Get("/api/v1/book", controllers.GetBooks)
	app.Get("/api/v1/book/:id", controllers.GetBook)
	app.Post("/api/v1/book", controllers.NewBook)
	app.Delete("/api/v1/book/:id", controllers.DeleteBook)
}

func initDatabase()  {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Database connected!!")

	database.DBConn.AutoMigrate(&models.Book{})
	fmt.Println("Database Migrated")
}

func main()  {
	app := fiber.New()

	initDatabase()

	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":3000")
}