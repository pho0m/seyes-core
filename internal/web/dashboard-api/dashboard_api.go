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

		r.Get("/", c.HealthCheck)
		r.Post("/notify", c.Notify)

		r.Get("/models/default", c.ReadModelFile)

		r.Route("/rooms", func(r chi.Router) {
			room := NewRoomController(sc)

			r.Get("/", room.IndexRoomHandler)
			r.Get("/{id}", room.GetRoomHandler)
			r.Post("/new", room.CreateRoomHandler)
			r.Put("/edit/{id}", room.UpdateRoomHandler)
			r.Delete("/delete/{id}", room.DeleteRoomHandler)
		})

		r.Route("/settings", func(r chi.Router) {
			room := NewSettingsController(sc)

			r.Get("/", room.GetSettingsHandler)
			r.Put("/edit/{id}", room.UpdateSettingsHandler)
		})

	})

}
