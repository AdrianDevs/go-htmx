package home

import (
	"fmt"
	"half.blue.gohtmx/util"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Load home page")
	util.Render(w, r, Page("Home"))
}
