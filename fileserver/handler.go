package fileserver

import (
	"github.com/go-chi/chi"
	"log/slog"
	"net/http"
	"strings"
)

type FileServer struct {
	Log  *slog.Logger
	Root http.Dir
}

func NewFileServer(log *slog.Logger, dir http.Dir) *FileServer {
	return &FileServer{Log: log, Root: http.Dir(dir)}
}

// Handler sets up a http.FileServer handler to serve static files-server from a http.FileSystem.
func (fs *FileServer) Handler(path string, r chi.Router) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		ctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(fs.Root))
		fs.ServeHTTP(w, r)
	})
}
