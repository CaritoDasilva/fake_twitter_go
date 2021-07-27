package routers

import (
	"encoding/json"
	"net/http"

	"github.com/fake_twitter/src/bd"
	"github.com/fake_twitter/src/models"
)

//Register crea usuarios en BD
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode((&t))
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El password tiene que tener un mínimo de 6 caracteres", 400)
		return
	}

	_, found, _ := bd.CheckUserExist(t.Email)

	if found {
		http.Error(w, "Usuario ya existe", 400)
		return
	}

	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al momento de registrar usuario"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Ocurrió un error al momento de registrar usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
