package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// This function sets up a new Fiber app, registers the Home route, and then sends a GET request
// to the root URL. It asserts that the response status code is 200 (OK).
//
// Running the tests: go test ./handlers
// Running all tests: go test ./...
//
// The ./... pattern tells Go to recursively find and run all tests in the current directory and all its subdirectories
func TestHome(t *testing.T) {
	app := fiber.New()
	app.Get("/", Home)

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}
