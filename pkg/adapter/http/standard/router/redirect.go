package standard_router

import (
	"net/http"
	"strings"

	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sub := len("/u/")
	shortenedURL := r.URL.String()[sub : sub+8]
	payload := dto.ShortenedDTO{ShortenedURL: shortenedURL}

	shortenerRepo := factory.NewShortenerRepository()
	statsRepo := factory.NewShortenerStatsRepository()
	ctl := controller.NewShortenerController(shortenerRepo, statsRepo)
	pre := ctl.Redirect(payload)
	res := pre.HTTP()

	if pre.Error == nil {
		if data, ok := res.Data.(dto.RedirectDTO); ok {
			http.Redirect(w, r, data.URL, res.StatusCode)
			return
		}
	}

	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res.JsonData)
}
