package main

import (
	"myGoApp/pkg/config"
	handerls "myGoApp/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handerls.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handerls.Repo.About))

	return mux

}
