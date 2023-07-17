package adapter

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	echo_router "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/http/echo/router"
)

func NewEchoServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/shorten", echo_router.Shorten)
	e.GET("/u/:shortenedURL", echo_router.Redirect)
	e.GET("/stats/:shortenedURL", echo_router.Stats)

	settings := config.GetSettings()
	port := fmt.Sprintf(":%s", settings.PORT)

	log.Println("Server listening on port", port)
	log.Println("Repository defined is:", settings.REPOSITORY_ADAPTER)

	e.Logger.Fatal(e.Start(port))
}