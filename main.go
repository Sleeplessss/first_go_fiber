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

	//	Grouping Routes
	api := app.Group("/api")
	v1 := api.Group("/v1") 

	v1.Get("/book", controllers.GetBooks)
	v1.Get("/book/:id", controllers.GetBook)
	v1.Post("/book", controllers.NewBook)
	v1.Delete("/book/:id", controllers.DeleteBook)
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