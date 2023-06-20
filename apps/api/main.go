package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	PORT := os.Getenv("PORT")
	port_number, err := strconv.Atoi(PORT)

	if err != nil {
		fmt.Println("Error PORT", err)
		return
	}

	app := fiber.New()

	app.Listen(fmt.Sprintf(":%d", port_number))
}
