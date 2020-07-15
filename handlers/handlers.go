package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/Neil-uli/tewto/middlew"
	"github.com/Neil-uli/tewto/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers handl*/
func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/register", middlew.CheckBD(routers.Register)),Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}