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
	r.HandleFunc("/api/personalities/new", controllers.New).Methods("POST")
	r.HandleFunc("/api/personalities/{id}/delete", controllers.Delete).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}/edit", controllers.Edit).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
