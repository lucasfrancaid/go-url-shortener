package main

import adapter "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/http/fiber"

func main() {
	adapter.NewFiberServer()
}
