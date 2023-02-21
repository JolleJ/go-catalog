package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jollej/item-catalog/pgm/controllers"
)

var RegisterItemRoutes = func(router *chi.Mux) {
	// RESTy routes for "articles" resource
	router.Route("/items", func(r chi.Router) {
		r.Get("/", controllers.GetItems)
		r.Get("/{id}", controllers.GetItem)
		r.Post("/", controllers.AddIitem)
		r.Delete("/{id}", controllers.DeleteItem)
	})

}
