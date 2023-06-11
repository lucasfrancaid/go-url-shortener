package main

import (
	adapter "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/http/standard"
)

func main() {
	adapter.NewHttpServer()
}
