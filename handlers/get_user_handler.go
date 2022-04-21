package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	token := r.Context().Value("token").(string)
	fmt.Println("user name:", vars["name"])
	fmt.Println("token auth = ", token)
	json.NewEncoder(w).Encode(struct {
		Result bool   `json:"result"`
		Name   string `json:"name"`
		Token  string `json:"token"`
	}{true, vars["name"], token})
}
