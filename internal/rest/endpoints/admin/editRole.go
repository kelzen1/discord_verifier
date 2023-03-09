package adminEndpoints

import (
	"encoding/json"
	"github.com/yoonaowo/discord_verifier/internal/database"
	"github.com/yoonaowo/discord_verifier/internal/models/rest"
	"github.com/yoonaowo/discord_verifier/internal/utils"
	"net/http"
)

func EditRole(w http.ResponseWriter, r *http.Request) {
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

	editRoleRequest, ok := utils.CastAndCompare[restModels.EditRoleReceiver](bodyData)
	if !ok {
		answer.Error = utils.ErrStructMismatch.Error()
		return
	}

	db := database.Get()
	err = db.EditRole(editRoleRequest.Name, editRoleRequest.Role)

	if err == nil {
		answer.Error = ""
	} else {
		answer.Error = err.Error()
	}

	answer.Success = err == nil

}
