package standard_router

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var payload dto.ShortenDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	ctl := controller.NewShortenerController()
	pre := ctl.Shorten(payload)
	res := pre.HTTP()

	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res.JsonData)
}
