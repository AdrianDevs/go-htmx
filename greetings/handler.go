package greet

import (
	"context"
	"fmt"
	"half.blue.gohtmx/util"
	"log/slog"
	"net/http"
	"strconv"
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
	HTMXReqHeader := r.Header.Get("HX-Request")
	hasHTMXReq, err := strconv.ParseBool(HTMXReqHeader)
	if err != nil {
		hasHTMXReq = false
	}

	var props ViewProps
	props.isHTMXReq = hasHTMXReq
	props.Title = "Greetings"
	props.Greeting, _ = h.Service.Hello(r.Context(), "bob")
	showView(w, r, props)

}

type ViewProps struct {
	isHTMXReq bool
	Title     string
	Greeting  string
}

func showView(w http.ResponseWriter, r *http.Request, props ViewProps) {
	fmt.Println("Load greetings page")
	fmt.Println("- isHTMXReq:", props.isHTMXReq)
	util.Render(w, r, Page(props.isHTMXReq, props.Title, props.Greeting))
}
