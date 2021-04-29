package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nelsongp/goSimpleWs/routers"
	"github.com/rs/cors"
)

func Manejadores(){
	router := mux.NewRouter()

	router.HandleFunc("/getBook", routers.BookGet).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "1122"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
