package home

import (
	"half.blue.gohtmx/lib"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, r, Show())
}
