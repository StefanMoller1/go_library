package router

import (
	"log"
	"net/http"
	"time"

	"github.com/StefanMoller1/go_library/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Manager struct {
	Log     *log.Logger
	Library app.Library
}

func (m *Manager) StartRouter() http.Handler {

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(-1))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))

	r.Mount("/api/v1", m.apiRouterV1())

	return r
}

func (m *Manager) apiRouterV1() http.Handler {
	r := chi.NewRouter()

	r.Route("/library", m.LibraryRouter)
	return r
}
