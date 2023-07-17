package main

import adapter "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/http/echo"

func main() {
	adapter.NewEchoServer()
}
