package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// DB health check
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		// 1) DB connection шалгах
		if err := db.Ping(); err != nil {
			http.Error(w, "DB NOT OK: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// 2) Users жагсаалт авах
		rows, err := db.Query(`SELECT id, name, email, created_at FROM users ORDER BY id`)
		if err != nil {
			http.Error(w, "DB NOT OK (query failed): "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
				http.Error(w, "DB NOT OK (scan failed): "+err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		// 3) JSON гаралт
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	// эндээс хойш products routes, auth гээд бүгдэд db дамжуулаад явж болно

	return r
}
