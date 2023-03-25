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
		r.Post("/notify", c.NotifyFromSeyesApp)
		// r.Get("/models/default", c.ReadModelFile)
		r.Route("/settings", func(r chi.Router) {
			room := NewSettingsController(sc)

			r.Get("/", room.GetSettingsHandler)
			r.Put("/edit/{id}", room.UpdateSettingsHandler)
		})

		r.Route("/rooms", func(r chi.Router) {
			roc := NewRoomController(sc)

			r.Get("/", roc.IndexRoomHandler)
			r.Get("/{id}", roc.GetRoomHandler)
			r.Post("/new", roc.CreateRoomHandler)
			r.Put("/edit/{id}", roc.UpdateRoomHandler)
			r.Delete("/delete/{id}", roc.DeleteRoomHandler)
		})

		r.Route("/reports", func(r chi.Router) {
			rec := NewReportController(sc)

			r.Get("/", rec.IndexReportHandler)
			r.Get("/{id}", rec.GetReportHandler)
			r.Post("/new", rec.CreateReportHandler)
			r.Put("/edit/{id}", rec.UpdateReportHandler)
			r.Delete("/delete/{id}", rec.DeleteReportHandler)
		})

		r.Route("/detect", func(r chi.Router) {
			dtc := NewDetectController(sc)

			r.Get("/{id}", dtc.GetDetectHandler)
			r.Post("/new", dtc.UpdateModelFileHandler)
		})
	})

}
