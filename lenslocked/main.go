package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/iamsushank/lenslocked/controllers"
	"path/filepath"

	// We want to add a new import similar to this one:
	"github.com/iamsushank/lenslocked/views"
)

func main() {
	r := chi.NewRouter()
	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))
	// Or inline everything and skip the `tpl` variable.
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))
	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))
}
