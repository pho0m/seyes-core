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

	// Shops Sections
	r.Route("/shops", func(r chi.Router) {
		c := NewShopsController(sc)
		// r.Use(a.MiddlewareAnalytic)
		r.Get("/", c.NewShop)
	})

}
