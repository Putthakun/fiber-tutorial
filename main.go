package main

import "github.com/gofiber/fiber/v2"

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
	app.Get("/books", func(c *fiber.Ctx) error  {
		return c.JSON(books)
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello wrold")
	})

	app.Listen((":8080"))
}
