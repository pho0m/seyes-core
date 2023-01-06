package dashboardAPI

import (
	"seyes-core/internal/service"

	"github.com/go-chi/chi"
)

// SetupRoutes setup route for dashboard-api
func SetupRoutes(sc *service.Container, r chi.Router) {

	// a := auth.NewApiMiddleware(sc)
	// // Auth Sections
	// r.Route("/login", func(r chi.Router) {
	// 	c := NewAuthController(sc)

	// 	r.Post("/", c.Login)
	// })

	// r.Route("/me", func(r chi.Router) {
	// 	c := NewAuthController(sc)

	// 	r.Get("/", c.Me)
	// 	r.Get("/session", c.GetSession)
	// 	r.Post("/session", c.CreateSession)
	// })

	// dashboard controller Sections
	r.Route("/api", func(r chi.Router) {
		c := NewDashboardController(sc)
		// r.Use(a.MiddlewareAnalytic)
		r.Post("/notify", c.Notify)
	})

}
