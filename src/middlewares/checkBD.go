package middlewares

import (
	"net/http"

	"github.com/fake_twitter/src/bd"
)

//CheckBD func middleware que chequea estado BD
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.CheckConnection() {
			http.Error(w, "Se ha perdido la conexión", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
