package api

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"myapp/db"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CheckCounter(w http.ResponseWriter, r *http.Request) {
	num := db.CounterProc()
	fmt.Fprintf(w, "Halaman telah dibuka sebanyak %+v kali.", num)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	users := db.Users{}
	
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := db.Insert(users); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	users, err := db.GetOne(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	if err := db.Remove(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	all, err := db.GetAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, all)
}

