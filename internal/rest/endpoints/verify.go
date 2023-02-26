package endpoints

import (
	"bytes"
	"encoding/json"
	"github.com/yoonaowo/discord_verifier/internal/database"
	restModels2 "github.com/yoonaowo/discord_verifier/internal/models/rest"
	utils2 "github.com/yoonaowo/discord_verifier/internal/utils"
	"io"
	"net/http"
)

func Verify(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	data := restModels2.VerifyAnswer{}

	defer func() {
		if !data.Success {
			marshaled, _ := json.Marshal(data)
			w.WriteHeader(400)
			_, _ = w.Write(marshaled)
		}
	}()

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		data.Error = err.Error()
		return
	}

	bodyData := make(map[string]any)
	err = json.Unmarshal(reqBody, &bodyData)

	if err != nil {
		data.Error = err.Error()
		return
	}

	if !utils2.CompareJSONToStruct(bodyData, restModels2.VerifyReceiver{}) {
		data.Error = utils2.ErrStructMismatch.Error()
		return
	}

	verifyRequest := &restModels2.VerifyReceiver{}
	_ = json.Unmarshal(reqBody, &verifyRequest)

	db := database.Get()
	data.Code, err = db.CreateOrGetCode(verifyRequest)

	if err != nil {
		return
	}

	data.Success = true

	marshaled, _ := json.Marshal(data)
	_, _ = w.Write(marshaled)
}
