package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"../../Projects/user-app-go/auth"
	"../../Projects/user-app-go/guests"
	"../../Projects/user-app-go/invites"
	"../../Projects/user-app-go/watchlists"
	"../../Projects/user-app-go/pushnotifications"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), 	// Set content-Type headers as application/json
		middleware.Logger,								// Log API request calls
		middleware.DefaultCompress,						// Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,						// Redirect slashes to no slash URL versions
		middleware.Recoverer,							// Recover from panics without crashing server
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/login", login.Routes())
		r.Mount("/api/guests", guests.Routes())
		r.Mount("/api/invites", invites.Routes())
		r.Mount("/api/watchlists", watchlists.Routes())
		r.Mount("/api/pushnotifications", pushnotifications.Routes())
	})

	return router
}

func main() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // Panic if there is an error
	}

	log.Fatal(http.ListenAndServe(":8000", router)) // Note, the port is usually retrieved from the environment
}