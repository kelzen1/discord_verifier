package rest

import (
	"Verifier/rest/endpoints"
	"Verifier/rest/middlewares"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Use(middlewares.CheckSignature)
	myRouter.HandleFunc("/verify", endpoints.Verify).Methods("POST")

	go func() {
		log.Fatalln(http.ListenAndServe(":80", myRouter))
	}()

}

func Init() {
	log.Println("[rest] start")
	handleRequests()
}
