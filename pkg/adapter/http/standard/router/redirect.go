package standard_router

import (
	"net/http"
	"strings"

	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

// Redirect godoc
//
//	@Summary		Redirect
//	@Description	Redirect a Shortened URL to Original URL
//	@Tags			Shortener
//	@Accept			json
//	@Produce		json
//	@Param			hashedURL	path	string	true	"Last block of Shortened URL, the value after /u/ part"
//	@Success		303
//	@Failure		404	{object}	presenter.ErrorResponseHTTP
//	@Failure		422	{object}	presenter.ErrorResponseHTTP
//	@Failure		500	{object}	presenter.ErrorResponseHTTP
//	@Router			/u/{hashedURL} [get]
func Redirect(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sub := len("/u/")
	shortenedURL := r.URL.String()[sub : sub+8]
	payload := dto.ShortenedDTO{ShortenedURL: shortenedURL}

	ctl := controller.NewShortenerController()
	pre := ctl.Redirect(payload)
	res := pre.HTTP()

	if pre.Error == nil {
		if data, ok := res.Data.(dto.RedirectDTO); ok {
			http.Redirect(w, r, data.URL, res.StatusCode)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	w.Write(res.JsonData)
}
