package adminEndpoints

import (
	"encoding/json"
	"github.com/yoonaowo/discord_verifier/internal/database"
	"github.com/yoonaowo/discord_verifier/internal/models/rest"
	"github.com/yoonaowo/discord_verifier/internal/utils"
	"net/http"
)

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	answer := restModels.AdminAnswer{
		Success: false,
		Error:   "unk",
	}

	defer func() {
		marshaled, _ := json.Marshal(answer)
		_, _ = w.Write(marshaled)
	}()

	bodyData, err := utils.ReadRequestBodyMap[string, any](r.Body)
	if err != nil {
		answer.Error = err.Error()
		return
	}

	deleteRoleRequest, ok := utils.CastAndCompare[restModels.DeleteRoleReceiver](bodyData)
	if !ok {
		answer.Error = utils.ErrStructMismatch.Error()
		return
	}

	db := database.Get()
	res := db.DeleteRole(deleteRoleRequest.Name)

	answer.Success = res == nil

	if res == nil {
		answer.Error = ""
	} else {
		answer.Error = res.Error()
	}

}
