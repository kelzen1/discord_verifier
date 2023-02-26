package rest

import (
	"github.com/go-chi/chi/v5"

	"github.com/yoonaowo/discord_verifier/internal/rest/endpoints"
	"github.com/yoonaowo/discord_verifier/internal/rest/middlewares"
	"github.com/yoonaowo/discord_verifier/internal/utils"

	"net/http"
)

func handleRequests() {

	myRouter := chi.NewRouter()

	myRouter.Use(middlewares.CheckSignature)
	myRouter.Post("/verify", endpoints.Verify)

	go func() {
		utils.Logger().Fatalln(http.ListenAndServe(":80", myRouter))
	}()

}

func Init() {
	utils.Logger().Println("starting rest")
	handleRequests()
}
