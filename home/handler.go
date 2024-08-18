package home

import (
	"half.blue.gohtmx/util"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	util.Render(w, r, Show())
}
