package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c *fiber.Ctx) error {
		type Query struct {
			Time    int64   `query:"t"`
			Action  string  `query:"a"`
			ClientX float64 `query:"x"`
			ClientY float64 `query:"y"`
		}
		query := &Query{}
		if err := c.QueryParser(query); err != nil {
			return fmt.Errorf("%w: %s", fiber.ErrBadRequest, err)
		}
		fmt.Println(query.Time, query.Action, query.ClientX, query.ClientY)
		return c.SendStatus(http.StatusNoContent)
	})

	// Start the server on port 3001
	log.Fatal(app.Listen(":3001"))
}
