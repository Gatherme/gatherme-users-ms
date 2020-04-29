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

var prefixPath = "/gatherme-users-ms"

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
func FindLikeByIDController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	like, err := connection.FindLikeByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, like)
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

// FindUserByEmailController - Encuentra un usuario por su email
func FindUserByEmailController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := connection.FindByEmail(params["email"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid email")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// FindUserByUsernameController - Encuentra un usuario por su username
func FindLikeByCategoryController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	like, err := connection.FindLikesByCategory(params["category"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category")
		return
	}
	respondWithJSON(w, http.StatusOK, like)
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

// CreatelikeController - Crear un gusto
func CreateLikeController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var like model.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	like.ID = bson.NewObjectId()
	if err := connection.InsertLike(like); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, like)
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

// UpdatelikeController - Actualiza un gusto
func UpdateLikeController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var like model.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.UpdateLike(like); err != nil {
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

// DeletelikeController - Borrar un usuario pot id
func DeleteLikeController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var likeID model.LikeID
	if err := json.NewDecoder(r.Body).Decode(&likeID); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.DeleteLike(likeID.ID); err != nil {
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
	r.HandleFunc(prefixPath+"/user-id/{id}", FindUserByIDController).Methods("GET")
	r.HandleFunc(prefixPath+"/user-email/{email}", FindUserByEmailController).Methods("GET")
	r.HandleFunc(prefixPath+"/user-username/{username}", FindUserByUsernameController).Methods("GET")

	r.HandleFunc(prefixPath+"/create-like", CreateLikeController).Methods("POST")
	r.HandleFunc(prefixPath+"/update-like", UpdateLikeController).Methods("PUT")
	r.HandleFunc(prefixPath+"/delete-like", DeleteLikeController).Methods("DELETE")
	r.HandleFunc(prefixPath+"/like-id/{id}", FindLikeByIDController).Methods("GET")
	r.HandleFunc(prefixPath+"/like-category/{category}", FindLikeByCategoryController).Methods("GET")

	log.Printf("Listening port 3000...")
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
