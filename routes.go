package main

import (
	"myGoApp/pkg/config"
	handerls "myGoApp/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handerls.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handerls.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handerls.Repo.Home)
	mux.Get("/about", handerls.Repo.About)

	return mux

}
