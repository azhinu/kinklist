package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/azhinu/kinklist/internal/storage"
)

func NewRouter(db *storage.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewHandler(db)

	r.Route("/api", func(r chi.Router) {
		r.Get("/kinklist/{id}", h.GetKinkList)
		r.Put("/kinklist/{id}", h.UpdateKinkList)
	})

	fs := http.FileServer(http.Dir("./frontend/dist"))
	r.Handle("/assets/*", fs)

	imgFS := http.StripPrefix("/img", http.FileServer(http.Dir("./frontend/public/img")))
	r.Handle("/img/*", imgFS)

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/dist/index.html")
	})

	return r
}
