package adapter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	standard_router "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/http/standard/router"
)

func NewHttpServer() {
	mux := http.NewServeMux()

	shortenHandler := http.HandlerFunc(standard_router.Shorten)
	mux.Handle("/shorten", shortenHandler)

	redirectHandler := http.HandlerFunc(standard_router.Redirect)
	mux.Handle("/u/", redirectHandler)

	statsHandler := http.HandlerFunc(standard_router.Stats)
	mux.Handle("/stats/", statsHandler)

	settings := config.GetSettings()
	port := fmt.Sprintf(":%s", settings.PORT)

	log.Println("Server listening on port", port)
	log.Println("Repository defined is:", settings.REPOSITORY_ADAPTER)

	_ = http.ListenAndServe(port, mux)
}
