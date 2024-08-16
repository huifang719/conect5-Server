package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)	

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", app.Home)
	mux.Get("/chat", app.ChatHandler)
	mux.Get("/game", app.GameHandler)
	mux.Get("/auth", app.Auth)
	return mux
}