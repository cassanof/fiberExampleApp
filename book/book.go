package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("Get a book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Adds a book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Deletes a book")
}
