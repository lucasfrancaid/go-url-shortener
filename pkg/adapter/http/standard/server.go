package adapter

import (
	"fmt"
	"log"
	"net/http"

	standard_router "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/http/standard/router"
)

func NewHttpServer() {
	mux := http.NewServeMux()

	shortenHandler := http.HandlerFunc(standard_router.Shorten)
	mux.Handle("/shorten", shortenHandler)

	port := fmt.Sprintf(":%s", "3333")
	log.Println("Server listening on port", port)
	_ = http.ListenAndServe(port, mux)
}
