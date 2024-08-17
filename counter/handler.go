package count

import (
	"context"
	"half.blue.gohtmx/lib"
	"log/slog"
	"net/http"
)

type CountService interface {
	Increment(ctx context.Context, it IncrementType, sessionID string) (counts Counts, err error)
	Get(ctx context.Context, sessionID string) (counts Counts, err error)
}

type IncHandler struct {
	Log          *slog.Logger
	CountService CountService
}

func NewHandler(log *slog.Logger, cs CountService) *IncHandler {
	return &IncHandler{
		Log:          log,
		CountService: cs,
	}
}

func (h *IncHandler) Get(w http.ResponseWriter, r *http.Request) {
	var props viewProps
	props.Counts.Global = 1
	props.Counts.Session = 2
	showView(w, r, props)
}

func (h *IncHandler) Post(w http.ResponseWriter, r *http.Request) {
	var props viewProps
	props.Counts.Global = 3
	props.Counts.Session = 4
	showView(w, r, props)
}

type viewProps struct {
	Counts Counts
}

func showView(w http.ResponseWriter, r *http.Request, props viewProps) {
	lib.Render(w, r, View(props.Counts.Global, props.Counts.Session))
}
