package standard_router

import (
	"net/http"
	"strings"

	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

func Stats(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sub := len("/stats/")
	shortenedURL := r.URL.String()[sub : sub+8]
	payload := dto.ShortenedDTO{ShortenedURL: shortenedURL}

	repo := factory.NewShortenerRepository()
	ctl := controller.NewShortenerController(repo)
	pre := ctl.Stats(payload)
	res := pre.HTTP()

	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res.JsonData)
}
