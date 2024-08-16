package handlers

import (
	"net/http"

	"half.blue.gohtmx/views/user"
)

type UserHandler struct {
	Service string
}

func New(service string) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("hi user"))
	// user.Show().Render(r.Context(), w)
	render(w, r, user.Show("bob"))
}
