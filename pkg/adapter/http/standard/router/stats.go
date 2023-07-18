package standard_router

import (
	"net/http"
	"strings"

	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

// Stats godoc
//
//	@Summary		Stats
//	@Description	Show statistics about a shortened URL
//	@Tags			Shortener
//	@Accept			json
//	@Produce		json
//	@Param			hashedURL	path		string	true	"Last block of Shortened URL, the value after /u/ part"
//	@Success		200			{object}	dto.ShortenerStatsDTO
//	@Failure		404			{object}	presenter.ErrorResponseHTTP
//	@Failure		422			{object}	presenter.ErrorResponseHTTP
//	@Failure		500			{object}	presenter.ErrorResponseHTTP
//	@Router			/stats/{hashedURL} [get]
func Stats(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sub := len("/stats/")
	shortenedURL := r.URL.String()[sub : sub+8]
	payload := dto.ShortenedDTO{ShortenedURL: shortenedURL}

	ctl := controller.NewShortenerController()
	pre := ctl.Stats(payload)
	res := pre.HTTP()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	w.Write(res.JsonData)
}
