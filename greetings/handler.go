package greet

import (
	"context"
	"half.blue.gohtmx/util"
	"log/slog"
	"net/http"
)

type ServiceInterface interface {
	Hello(ctx context.Context, name string) (string, error)
	Hellos(ctx context.Context, names []string) (map[string]string, error)
}

type Handler struct {
	Log     *slog.Logger
	Service ServiceInterface
}

func NewHandler(log *slog.Logger, s ServiceInterface) *Handler {
	return &Handler{
		Log:     log,
		Service: s,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.get(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 Method Not Allowed. Only GET is allowed."))
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	var props ViewProps
	props.Greeting, _ = h.Service.Hello(r.Context(), "bob")
	showView(w, r, props)

}

type ViewProps struct {
	Greeting string
}

func showView(w http.ResponseWriter, r *http.Request, props ViewProps) {
	util.Render(w, r, View(props.Greeting))
}
