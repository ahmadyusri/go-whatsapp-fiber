package http

import (
	"os"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func NewSwaggerHandler(app *fiber.App) {
	BASIC_AUTH_USERNAME := os.Getenv("BASIC_AUTH_USERNAME")
	BASIC_AUTH_PASSWORD := os.Getenv("BASIC_AUTH_PASSWORD")

	if BASIC_AUTH_USERNAME != "" && BASIC_AUTH_PASSWORD != "" {
		// Basic Auth Provide a minimal config
		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				BASIC_AUTH_USERNAME: BASIC_AUTH_PASSWORD,
			},
		}))
	}

	// Routes for GET method:

	//default
	app.Get("/swagger/*", swagger.Handler) // get one user by ID

	/*app.Get("/swagger/*", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})*/

	//custom
	/*route.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: "http://example.com/doc.json",
		DeepLinking: false,
	}))*/
}
