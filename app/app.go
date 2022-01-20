package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bayusm0506/hcf-or-gcd/config"
	"github.com/gorilla/mux"
)

// App has router
type App struct {
	Router *mux.Router
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

// SetRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling
}

// Post wraps the routers for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Run the app on it's router
func (a *App) Run(config *config.Config) {
	port := fmt.Sprintf(":%s",
		config.APPS.Port,
	)

	log.Println("Listening on port : " + port)

	log.Fatal(http.ListenAndServe(port, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
