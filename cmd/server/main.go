package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/kitmacled/go-rest-api/internal/transport/http"
)

// App - the struct that contains things like pointers
// to database connections
type App struct{}

// Run - sets up our app
func (app *App) Run() error {
	fmt.Println("Setting up our app")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8081", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
