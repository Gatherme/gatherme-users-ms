package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Gatherme/gatherme-users-ms/connection"
	"github.com/Gatherme/gatherme-users-ms/model"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var prefixPath = "/api/users"

// FindUserByIDController - Encuentra un usuario por su ID
func FindUserByIDController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := connection.FindUserByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// FindUserByIDController - Encuentra un usuario por su ID
func FindPleasureByIDController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pleasure, err := connection.FindPleasureByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, pleasure)
}

// FindUserByUsernameController - Encuentra un usuario por su username
func FindUserByUsernameController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := connection.FindByUsername(params["username"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid username")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// FindUserByUsernameController - Encuentra un usuario por su username
func FindPleasureByCategoryController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pleasure, err := connection.FindPleasuresByCategory(params["category"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category")
		return
	}
	respondWithJSON(w, http.StatusOK, pleasure)
}

// CreateUserController - Crear un usuario
func CreateUserController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	user.ID = bson.NewObjectId()
	if err := connection.InsertUser(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// CreatePleasureController - Crear un gusto
func CreatePleasureController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var pleasure model.Pleasure
	if err := json.NewDecoder(r.Body).Decode(&pleasure); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	pleasure.ID = bson.NewObjectId()
	if err := connection.InsertPleasure(pleasure); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, pleasure)
}

// UpdateUserController - Actualiza un usuario
func UpdateUserController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.UpdateUser(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// UpdatePleasureController - Actualiza un gusto
func UpdatePleasureController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var pleasure model.Pleasure
	if err := json.NewDecoder(r.Body).Decode(&pleasure); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.UpdatePleasure(pleasure); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteUserController - Borrar un usuario pot id
func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userID model.UserID
	if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.DeleteUser(userID.ID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeletePleasureController - Borrar un usuario pot id
func DeletePleasureController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var pleasureID model.PleasureID
	if err := json.NewDecoder(r.Body).Decode(&pleasureID); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.DeletePleasure(pleasureID.ID); err != nil {
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

	r.HandleFunc(prefixPath+"/create-pleasure", CreatePleasureController).Methods("POST")
	r.HandleFunc(prefixPath+"/update-pleasure", UpdatePleasureController).Methods("PUT")
	r.HandleFunc(prefixPath+"/delete-pleasure", DeletePleasureController).Methods("DELETE")
	r.HandleFunc(prefixPath+"/pleasure-id/{id}", FindPleasureByIDController).Methods("GET")
	r.HandleFunc(prefixPath+"/pleasure-category/{category}", FindPleasureByCategoryController).Methods("GET")

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
