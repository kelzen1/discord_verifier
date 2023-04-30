package rest

import (
	"github.com/go-chi/chi/v5"
	adminEndpoints "github.com/yoonaowo/discord_verifier/internal/rest/endpoints/admin"
	userEndpoints "github.com/yoonaowo/discord_verifier/internal/rest/endpoints/user"
	"github.com/yoonaowo/discord_verifier/internal/rest/middlewares"
	"github.com/yoonaowo/discord_verifier/internal/utils"

	"net/http"
)

func userRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares.CheckSignature)

	router.Post("/check_verification", userEndpoints.Verify)

	return router
}

func adminRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares.CheckAdminToken)

	router.Post("/role_editor", adminEndpoints.EditRole)
	router.Post("/role_delete", adminEndpoints.DeleteRole)
	router.Get("/role_list", adminEndpoints.ListRoles)

	return router
}

func handleRequests() {

	mainRouter := chi.NewRouter()
	mainRouter.Mount("/", userRouter())
	mainRouter.Mount("/superpower", adminRouter())

	go utils.Logger().Fatalln(http.ListenAndServe(":"+utils.FlagWebPort, mainRouter))
}

func Init() {
	utils.Logger().Println("starting rest")
	handleRequests()
}
