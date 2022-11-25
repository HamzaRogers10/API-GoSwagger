package routes

import (
	"API-GoSwagger/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func InitializeRouters() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello/{name}", handlers.Index).Methods("GET")
	router.HandleFunc("/swagger.json", handlers.Swagger).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}
