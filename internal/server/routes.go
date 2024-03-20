package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// The routes() method returns a servemux containing our application routes.
func (app *Application) routes() *chi.Mux {

	r := chi.NewRouter()
	// r.Use(middleware.Logger)

	r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// r.Use(middleware.URLFormat)
	// r.Use(render.SetContentType(render.ContentTypeJSON))

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })
	r.Get("/", app.home)
	r.Get("/snippet/view", app.snippetView)
	r.Post("/snippet/create", app.snippetCreate)

	// todo controller

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/todos", app.getAllTodos)
		r.Get("/todo/{id}", app.getTodoById)
		r.Post("/todo", app.createTodo)
		r.Put("/todo/edit", app.editTodo)
		r.Put("/todo/checked/{id}", app.updateTodoCheck)
		r.Delete("/todo/{id}", app.deleteTodo)
	})

	return r
}
