package api

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vitormuuniz/winestore-go/api/controllers"
	"github.com/vitormuuniz/winestore-go/api/database"
	"github.com/vitormuuniz/winestore-go/api/repositories"
	"github.com/vitormuuniz/winestore-go/api/routes"
)

var (
	port = flag.Int("p", 5000, "set port")
)

func Run() {
	flag.Parse()

	db := database.Connect()
	if db != nil {
		defer db.Close()
	}
	
	wsRepository := repositories.NewWineStoreRepository(db)

	wsController := controllers.NewWineStoreController(wsRepository)

	wsRoutes := routes.NewWineStoreRoutes(wsController)

	router := mux.NewRouter().StrictSlash(true)

	routes.Install(router, wsRoutes)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested", "Location"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Printf("Server running on port %d \n", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.CORS(headers, methods, origins)(router)))
}
