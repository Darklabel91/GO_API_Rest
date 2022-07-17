package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Darklabel91/GO_API_Rest/database"
	models "github.com/Darklabel91/GO_API_Rest/models"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ReturnAllPersonalities(w http.ResponseWriter, r *http.Request) {

	var personalities []models.Personality
	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func ReturnPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)

	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	var personality models.Personality
	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(personality)
}
