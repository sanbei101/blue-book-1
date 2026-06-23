package api

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/sanbei101/blue-book/internal/db"
	"github.com/sanbei101/blue-book/internal/pkg/jwt"
)

func RegisterRoutes(store *db.Store) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	postHandler := NewPostHandler(store)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/posts", postHandler.ListFeed)
		r.Get("/posts/{id}", postHandler.GetByID)
		r.Get("/posts/user/{userID}", postHandler.ListByUser)

		r.Group(func(r chi.Router) {
			r.Use(jwt.AuthMiddleware)
			r.Post("/posts", postHandler.Create)
			r.Delete("/posts/{id}", postHandler.Delete)
		})
	})

	return r
}
