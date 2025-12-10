package http

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// DB health check
	r.Get("/db-health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(w, "DB NOT OK: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("DB OK"))
	})

	// эндээс хойш products routes, auth гээд бүгдэд db дамжуулаад явж болно

	return r
}
