package routes

import (
	"apiRest/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.ReturnPersonality).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
