package util

import (
	"fmt"
	"net/http"

	"gitlab.com/go-htmx/go-htmx/pkg/htmx"
)

func CheckIfHtmxMiddlerware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Incoming HTTP request", r.Method, r.URL.String())

		ctx := htmx.NewContext(r)
		// ctx.HxTarget = "#main"
		if ctx.IsHtmxRequest {
			fmt.Println("- isHTMXReq: True")
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("- isHTMXReq: False")
			next.ServeHTTP(w, r)
		}
	})
}
