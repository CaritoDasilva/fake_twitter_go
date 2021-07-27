package routers

import (
	"encoding/json"
	"net/http"

	"github.com/fake_twitter/src/bd"
)

func ShowProfile(w http.ResponseWriter, r http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	profile, err := bd.FindProfile(ID)

	if err != nil {
		http.Error(w, "Debe enviar el parámetro ID"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)

}
