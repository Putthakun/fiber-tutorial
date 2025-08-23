package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Book struct {
    ID 		int 	`json:"id"`
	Title 	string 	`json:"titel"`
	Auhtor 	string 	`json:"author"`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Obuto", Auhtor: "Putthakun"})
	books = append(books, Book{ID: 2, Title: "JK", Auhtor: "Putthakun"})

	// get all data
	app.Get("/books", getBooks)

	//get by id
	app.Get("/books/:id", getBook)

	// localhost
	app.Listen((":8080"))
}

func getBooks(c *fiber.Ctx) error {
	return c.SendString("Hello wrold")
}

func getBook(c *fiber.Ctx) error {
	bookId, error := strconv.Atoi(c.Params("id"))

	if error != nil {
		return c.Status(fiber.StatusBadRequest).SendString(error.Error()) // return status 400 and return error
	}

	for _, book := range(books) {
		if book.ID == bookId {
			return c.JSON(book) // return book by id json type
		}
	}

	return c.SendStatus(fiber.StatusNotFound) // return status 404

	// custom error
	// return c.Status(fiber.StatusNotFound).SendString("Book not found") return status and return error
}