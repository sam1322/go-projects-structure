package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })
	r.Get("/", app.home)
	r.Get("/snippet/view", app.snippetView)
	r.Post("/snippet/create", app.snippetCreate)
	return r
}
