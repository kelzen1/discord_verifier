package userEndpoints

import (
	"github.com/yoonaowo/discord_verifier/internal/database"
	"github.com/yoonaowo/discord_verifier/internal/models/rest"
	"github.com/yoonaowo/discord_verifier/internal/utils"

	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Verify(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	data := restModels.VerifyAnswer{}

	defer func() {
		if !data.Success {
			marshaled, _ := json.Marshal(data)
			w.WriteHeader(400)
			_, _ = w.Write(marshaled)
		}
	}()

	bodyData, err := utils.ReadRequestBodyMap[string, any](r.Body)
	if err != nil {
		data.Error = err.Error()
		return
	}

	verifyRequest, ok := utils.CastAndCompare[restModels.VerifyReceiver](bodyData)
	if !ok {
		data.Error = utils.ErrStructMismatch.Error()
		return
	}

	db := database.Get()

	_, err = db.GetRoleID(verifyRequest.Role)

	if err != nil {
		data.Error = utils.ErrRoleNotFound.Error()
		return
	}

	data.Code, err = db.CreateOrGetCode(verifyRequest)

	if err != nil {
		return
	}

	data.Success = true

	marshaled, _ := json.Marshal(data)
	_, _ = w.Write(marshaled)
}
