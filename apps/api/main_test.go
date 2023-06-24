package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func testStatus(t *testing.T, app *fiber.App, url, method string) {

	t.Helper()

	req := httptest.NewRequest(method, url, nil)

	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 404, res.StatusCode, "Status code")
}

func TestAppMethods(t *testing.T) {
	app := fiber.New()

	testStatus(t, app, "/", "GET")
	testStatus(t, app, "/create", "POST")
}
