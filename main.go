package main

import (
	"encoding/json"
	"log"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/Gatherme/gatherme-users-ms/connection"
	"github.com/Gatherme/gatherme-users-ms/model"
	"github.com/gorilla/mux"

)

var prefixPath = "/api/users"

// FindUserByIDController - Encuentra un usuario por su ID
func FindUserByIDController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := connection.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// FindUserByUsernameController - Encuentra un usuario por su username
func FindUserByUsernameController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user,err := connection.FindByUsername(params["username"]) 
	
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid username")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// CreateUserController - Crear un usuario
func CreateUserController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shopping model.User
	if err := json.NewDecoder(r.Body).Decode(&shopping); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	shopping.ID = bson.NewObjectId()
	if err := connection.Insert(shopping); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, shopping)
}

// UpdateUserController - Actualiza un usuario 
func UpdateUserController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shopping model.User
	if err := json.NewDecoder(r.Body).Decode(&shopping); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.Update(shopping); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteUserController - Borrar un usuario pot id
func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shoppingID model.UserID
	if err := json.NewDecoder(r.Body).Decode(&shoppingID); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.Delete(shoppingID.ID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc(prefixPath+"/create-user", CreateUserController).Methods("POST")
	r.HandleFunc(prefixPath+"/update-user", UpdateUserController).Methods("PUT")
	r.HandleFunc(prefixPath+"/delete-user", DeleteUserController).Methods("DELETE")
	r.HandleFunc(prefixPath+"/by-id/{id}", FindUserByIDController).Methods("GET")
	r.HandleFunc(prefixPath+"/by-username/{username}", FindUserByUsernameController).Methods("GET")

	log.Printf("Listening...")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}