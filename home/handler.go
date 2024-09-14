package home

import (
	"fmt"
	"net/http"
	"strconv"

	"half.blue.gohtmx/util"
)

func Index(w http.ResponseWriter, r *http.Request) {
	HTMXReqHeader := r.Header.Get("HX-Request")
	hasHTMXReq, err := strconv.ParseBool(HTMXReqHeader)
	if err != nil {
		hasHTMXReq = false
	}

	var props ViewProps
	props.isHTMXReq = hasHTMXReq
	props.Title = "Home"
	showView(w, r, props)
}

type ViewProps struct {
	isHTMXReq bool
	Title     string
}

func showView(w http.ResponseWriter, r *http.Request, props ViewProps) {
	fmt.Println("Show home page")
	util.Render(w, r, Content(props.isHTMXReq, props.Title))

}
