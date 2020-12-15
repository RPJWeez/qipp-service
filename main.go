package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"qip.io/q/pkg/qipp/database"
	"qip.io/q/pkg/qipp/handler"
	"qip.io/q/pkg/qipp/middleware"
)

type Response struct {
	Message string `json:"message"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

func main() {
	database.InitDb()
	defer database.DbConn.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Using port " + port)

	router := mux.NewRouter()

	jm := middleware.GetAuthMw()
	um := middleware.UserIdMiddleware

	router.Handle("/q/{id}", negroni.New(
		negroni.HandlerFunc(jm.HandlerWithNext),
		negroni.HandlerFunc(um),
		negroni.Wrap(http.HandlerFunc(handler.GetQipp)))).Methods("GET")

	router.Handle("/q", negroni.New(
		negroni.HandlerFunc(jm.HandlerWithNext),
		negroni.HandlerFunc(um),
		negroni.Wrap(http.HandlerFunc(handler.PostQipp)))).Methods("POST")

	router.Handle("/q", negroni.New(
		negroni.HandlerFunc(jm.HandlerWithNext),
		negroni.HandlerFunc(um),
		negroni.Wrap(http.HandlerFunc(handler.GetQipps)))).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
