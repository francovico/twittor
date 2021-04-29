package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/francovico/twittor/middlew"
	"github.com/francovico/twittor/routers"
)

// Manejadores, seteo mi puerto, el Handler y pongo a escuchar al servidor.
func Manejadores() {
	println("Inicializando handlers")
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.Login))).Methods("POST")

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
