package standard_router

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

// Shorten godoc
//
//	@Summary		Shorten
//	@Description	Shorten an URL
//	@Tags			Shortener
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		dto.ShortenDTO	true	"URL to Shorten"
//	@Success		200		{object}	dto.ShortenedDTO
//	@Failure		422		{object}	presenter.ErrorResponseHTTP
//	@Failure		500		{object}	presenter.ErrorResponseHTTP
//	@Router			/shorten [post]
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	w.Write(res.JsonData)
}
