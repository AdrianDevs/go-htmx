package util

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"gitlab.com/go-htmx/go-htmx/pkg/htmx"
	"half.blue.gohtmx/components"
)

func isLocalhost(r *http.Request) bool {
	return strings.Contains(r.Host, "localhost")
}

func IsHTMXRequest(r *http.Request) bool {
	ctx := htmx.NewContext(r)
	return ctx.IsHtmxRequest
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) {
	if IsHTMXRequest(r) {
		err := c.Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		// If it's not an HTMX request, render the full page.
		// This is useful for when the user navigates to the page directly.
		err := components.HTMLWrapper(c).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
