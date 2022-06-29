package controllers

import (
	"apiRest/models"
	"apiRest/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "HomePage")

}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := repository.FindAll()
	json.NewEncoder(w).Encode(p)
}

func ReturnPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	p := repository.FindId(id)

	json.NewEncoder(w).Encode(p)

}

func New(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	repository.Create(&personality)
	json.NewEncoder(w).Encode(personality)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	p := repository.Delete(id)

	json.NewEncoder(w).Encode(p)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	var personality models.Personality
	personality.Id = id
	json.NewDecoder(r.Body).Decode(&personality)
	repository.Update(&personality)
	json.NewEncoder(w).Encode(personality)
}
