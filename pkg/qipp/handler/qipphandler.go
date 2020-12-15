package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"qip.io/q/pkg/qipp/database"
	"qip.io/q/pkg/qipp/model"
)

func PostQipp(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	if db == nil {
		fmt.Println("DB no bueno")
	}

	user := r.Header.Get("x-user-id")

	var qipp model.Qipp
	json.NewDecoder(r.Body).Decode(&qipp)

	qipp.UserId = user

	db.Create(&qipp)

	json.NewEncoder(w).Encode(qipp)
	return
}

func GetQipp(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn

	params := mux.Vars(r)
	id := params["id"]

	var qipp model.Qipp
	db.Find(&qipp, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(qipp)
	return
}

func GetQipps(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	params := mux.Vars(r)
	user := params["u"]

	if len(user) == 0 {
		user = r.Header.Get("x-user-id")
	}

	var qipps []model.Qipp
	db.Find(&qipps, "user_id = ?", user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(qipps)
}
