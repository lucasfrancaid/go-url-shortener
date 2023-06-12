package standard_router

import (
	"encoding/json"
	"net/http"
	"strings"

	adapter "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
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

	repo := adapter.NewShortenerRepositoryInMemory()
	ctl := controller.NewShortenerController(&repo)
	pre := ctl.Redirect(payload)
	res := pre.HTTP()

	if pre.Error == nil {
		var readDTO dto.RedirectDTO
		err := json.Unmarshal(res.Data, &readDTO)
		if err == nil {
			http.Redirect(w, r, readDTO.URL, res.StatusCode)
			return
		}
	}

	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res.Data)
}
