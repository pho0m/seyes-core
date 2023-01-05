package web

import (
	"fmt"
	"net/http"
	"seyes-core/internal/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"seyes-core/internal/web/common"
	dashboardAPI "seyes-core/internal/web/dashboard-api"
)

// Server defines a Web Server
type Server struct {
	Port   string
	Router *chi.Mux
}

// NewServer creates a new Server
func NewServer(sc *service.Container, port string) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Token", "Liff-Id", "Line-Token"},
		AllowCredentials: true,
	}))

	r.Route("/", func(r chi.Router) {
		dashboardAPI.SetupRoutes(sc, r)
	})

	return &Server{
		Port:   port,
		Router: r,
	}
}

// // NewApiMiddleware creates a new NewApiMiddleware
// func NewApiMiddleware() map[string]string {
// 	username := os.Getenv("BASIC_AUTH_USERNAME")
// 	password := os.Getenv("BASIC_AUTH_PASSWORD")

// 	cred := map[string]string{
// 		username: password,
// 	}

// 	return cred
// }
// RegisterHandler add handler to router
func (s *Server) RegisterHandler(h common.Handler) {
	h.Register(s.Router)
}

// Start starts a web server
func (s *Server) Start(sc *service.Container) {

	if err := http.ListenAndServe(":"+s.Port, s.Router); err != nil {
		panic(fmt.Sprintf("cannot listen to port (%s): %s", s.Port, err.Error()))
	}
}
