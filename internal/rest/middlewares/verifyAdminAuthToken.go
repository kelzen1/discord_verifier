package middlewares

import (
	"encoding/json"
	restModels "github.com/yoonaowo/discord_verifier/internal/models/rest"
	"github.com/yoonaowo/discord_verifier/internal/utils"
	"net/http"
)

func CheckAdminToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("admin_key") != utils.FlagAdminPassword {
			marshaled, _ := json.Marshal(restModels.ErrorAnswer{
				Success: false,
				Error:   "wrong_admin_key",
			})
			w.WriteHeader(400)
			_, _ = w.Write(marshaled)
			return
		}

		next.ServeHTTP(w, r)
	})
}
