package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/fake_twitter/src/middlewares"
	"github.com/fake_twitter/src/routers"
)

//Handlers escucha el servidor
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/api/registro", middlewares.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/api/login", middlewares.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/api/perfil", middlewares.CheckBD(middlewares.JwtValidation(routers.ShowProfile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
