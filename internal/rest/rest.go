package rest

import (
	"github.com/gorilla/mux"
	"github.com/yoonaowo/discord_verifier/internal/rest/endpoints"
	"github.com/yoonaowo/discord_verifier/internal/rest/middlewares"
	"github.com/yoonaowo/discord_verifier/internal/utils"
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
