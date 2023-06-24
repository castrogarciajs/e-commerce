package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/sebastian009w/e-commerce/pkg/routes"
)

func main() {

	app := fiber.New()

	app.Get("/", routes.Home)
	app.Listen(":3000")
	fmt.Println("Server on port: 3000")
}
