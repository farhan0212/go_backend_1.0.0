package main

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func mockSetupRoutes(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}

func TestFiberAppInitialization(t *testing.T) {
	app := fiber.New()
	assert.NotNil(t, app, "Fiber app should be initialized")
}

func TestPingRoute(t *testing.T) {
	app := fiber.New()
	mockSetupRoutes(app)

	req := &fiber.Ctx{}
	resp, err := app.Test(req.Request(), -1)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCORSConfig(t *testing.T) {
	app := fiber.New()
	// Add CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		return c.Next()
	})
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, "*", resp.Header.Get("Access-Control-Allow-Origin"))
}