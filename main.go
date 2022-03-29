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

	http.HandleFunc("/", handerls.Home)
	http.HandleFunc("/about", handerls.About)
	// http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting appliaction on port %s", portNumber))

	http.ListenAndServe(portNumber, nil)
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
