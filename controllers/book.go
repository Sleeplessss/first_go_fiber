package controllers

import (
	"github.com/Sleeplessss/first_go_fiber/database"
	"github.com/Sleeplessss/first_go_fiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []models.Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book models.Book
	db.Find(&book, id)
	return c.JSON(&book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book models.Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(400).SendString("Book not found.")
	}
	db.Delete(&book)
	return c.JSON(&book)
}
