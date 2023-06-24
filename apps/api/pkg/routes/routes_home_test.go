package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestHome(t *testing.T) {

	t.Parallel()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Home")
	})

	res, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/", nil))

	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 200, res.StatusCode, "Status Code")
}
