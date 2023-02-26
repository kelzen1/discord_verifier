package rest

import (
	"Verifier/rest/endpoints"
	"Verifier/rest/middlewares"
	"Verifier/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Use(middlewares.CheckSignature)
	myRouter.HandleFunc("/verify", endpoints.Verify).Methods("POST")

	go func() {
		utils.Logger().Fatalln(http.ListenAndServe(":80", myRouter))
	}()

}

func Init() {
	utils.Logger().Println("starting rest")
	handleRequests()
}
