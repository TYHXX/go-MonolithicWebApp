package main

import (
	"errors"
	"fmt"
	"log"
	"myGoApp/pkg/config"
	handerls "myGoApp/pkg/handlers"
	"myGoApp/pkg/render"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can't create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handerls.NewRepo(&app)
	handerls.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handerls.Repo.Home)
	// http.HandleFunc("/about", handerls.Repo.About)
	// http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting appliaction on port %s", portNumber))

	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal((err))
}

func addValue(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32 = 100.0
	var y float32 = 0.0

	f, err := devideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))

}

func devideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cant devide by zero")
		return 0, err
	}
	result := x / y

	return result, nil
}
