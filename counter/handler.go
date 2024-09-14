package count

import (
	"context"
	"log/slog"
	"net/http"

	"half.blue.gohtmx/util"
)

type ServiceInterface interface {
	Increment(ctx context.Context, it IncrementType, sessionID string) (counts Counts, err error)
	Get(ctx context.Context, sessionID string) (counts Counts, err error)
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

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	var props viewProps
	props.Counts.Global = 1
	props.Counts.Session = 2
	showView(w, r, props)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var props viewProps
	props.Counts.Global = 3
	props.Counts.Session = 4
	showView(w, r, props)
}

type viewProps struct {
	Counts Counts
}

func showView(w http.ResponseWriter, r *http.Request, props viewProps) {
	util.Render(w, r, View(props.Counts.Global, props.Counts.Session))
}
