package adapter

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lucasfrancaid/go-url-shortener/docs"
	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	standard_middleware "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/http/standard/middleware"
	standard_router "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/http/standard/router"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			URL Shortener API
//	@version		1.0
//	@description	Go URL Shortener implemented using Clean Architecture with Echo and Fiber as HTTP Adapters.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:3000
//	@BasePath		/
func NewHttpServer() {
	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	shortenHandler := http.HandlerFunc(standard_router.Shorten)
	mux.Handle("/shorten", standard_middleware.LoggerMiddleware(shortenHandler))

	redirectHandler := http.HandlerFunc(standard_router.Redirect)
	mux.Handle("/u/", standard_middleware.LoggerMiddleware(redirectHandler))

	statsHandler := http.HandlerFunc(standard_router.Stats)
	mux.Handle("/stats/", standard_middleware.LoggerMiddleware(statsHandler))

	settings := config.GetSettings()
	port := fmt.Sprintf(":%s", settings.PORT)

	log.Println("Server listening on port", port)
	log.Println("Repository defined is:", settings.REPOSITORY_ADAPTER)

	_ = http.ListenAndServe(port, mux)
}
