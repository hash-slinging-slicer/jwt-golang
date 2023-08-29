package main

import (
	authcontroller "jwt/controllers/authcontroller"
	"jwt/controllers/produkcontroller"
	"jwt/middlewares"
	"jwt/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.KonekDB()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	// SUB ROUTER
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/produk", produkcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleWare)

	log.Fatal(http.ListenAndServe(":8080", r))
}
