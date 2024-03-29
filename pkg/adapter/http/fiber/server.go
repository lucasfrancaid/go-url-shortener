package adapter

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "github.com/lucasfrancaid/go-url-shortener/docs"
	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	fiber_router "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/http/fiber/router"
)

func NewFiberServer() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/shorten", fiber_router.Shorten)
	app.Get("/u/:hashedURL", fiber_router.Redirect)
	app.Get("/stats/:hashedURL", fiber_router.Stats)

	settings := config.GetSettings()
	port := fmt.Sprintf(":%s", settings.PORT)

	log.Println("Server listening on port", port)
	log.Println("Repository defined is:", settings.REPOSITORY_ADAPTER)

	log.Fatal(app.Listen(port))
}
