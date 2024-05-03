package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/itsrojasleon/messaging-app/auth/internal/auth"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", auth.SignupHandler)
		r.Post("/signin", auth.SigninHandler)
		r.Get("/currentuser", auth.CurrentUserHandler)
	})

	http.ListenAndServe(":8080", r)
}
