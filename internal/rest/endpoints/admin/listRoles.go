package adminEndpoints

import (
	"encoding/json"
	"github.com/yoonaowo/discord_verifier/internal/database"
	"github.com/yoonaowo/discord_verifier/internal/models/rest"
	"net/http"
)

func ListRoles(w http.ResponseWriter, _ *http.Request) {
	answer := restModels.AdminAnswer{
		Success: false,
		Error:   "unk",
	}

	defer func() {
		marshaled, _ := json.Marshal(answer)
		_, _ = w.Write(marshaled)
	}()

	db := database.Get()
	roles := db.ListRoles()
	mappedRoles := make(map[string]string, len(roles))

	for _, v := range roles {
		mappedRoles[v.Name] = v.Role
	}

	answer.Error = ""
	answer.Success = true
	answer.Data = mappedRoles
}
