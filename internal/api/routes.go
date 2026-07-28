package api

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/sanbei101/blue-book/internal/db"
	"github.com/sanbei101/blue-book/internal/pkg/jwt"
)

// OpenAPI 定义
//
//	@title							小蓝书 API
//	@version						1.0
//	@description					小蓝书后端接口文档
//	@servers.url					http://localhost:8080/api/v1
//	@servers.description			本地开发环境
//	@securitydefinitions.bearerauth	BearerAuth
func RegisterRoutes(store *db.Store) *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	postHandler := NewPostHandler(store)
	userHandler := NewUserHandler(store)
	commentHandler := NewCommentHandler(store)
	likeHandler := NewLikeHandler(store)
	followHandler := NewFollowHandler(store)

	r.Route("/api/v1", func(r chi.Router) {
		// 公开路由
		r.Post("/users/register", userHandler.Register)
		r.Post("/users/login", userHandler.Login)
		r.Get("/users/{id}", userHandler.GetProfile)
		r.Get("/posts", postHandler.ListFeed)
		r.Get("/posts/{id}", postHandler.GetByID)
		r.Get("/posts/user/{userID}", postHandler.ListByUser)
		r.Get("/posts/{id}/comments", commentHandler.ListByPost)
		r.Get("/users/{id}/followers", followHandler.ListFollowers)
		r.Get("/users/{id}/following", followHandler.ListFollowing)

		// 需要认证的路由
		r.Group(func(r chi.Router) {
			r.Use(jwt.AuthMiddleware)
			r.Put("/users/profile", userHandler.UpdateProfile)
			r.Post("/posts", postHandler.Create)
			r.Delete("/posts/{id}", postHandler.Delete)
			r.Post("/comments", commentHandler.Create)
			r.Post("/likes", likeHandler.Toggle)
			r.Post("/users/{id}/follow", followHandler.Follow)
			r.Delete("/users/{id}/follow", followHandler.Unfollow)
		})
	})

	return r
}
