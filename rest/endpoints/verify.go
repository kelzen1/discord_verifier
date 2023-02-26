package endpoints

import (
	"Verifier/database/actions"
	restModels "Verifier/models/rest"
	"Verifier/utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Verify(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	r.Body.Close() //  must close
	r.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	data := restModels.VerifyAnswer{}

	defer func() {
		if !data.Success {
			marshaled, _ := json.Marshal(data)
			w.WriteHeader(400)
			w.Write(marshaled)
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

	if !utils.CompareJSONToStruct(bodyData, restModels.VerifyReceiver{}) {
		data.Error = utils.ErrStructMismatch.Error()
		return
	}

	verifyRequest := &restModels.VerifyReceiver{}
	_ = json.Unmarshal(reqBody, &verifyRequest)

	data.Code, err = actions.CreateOrGetCode(verifyRequest)

	if err != nil {
		return
	}

	data.Success = true

	marshaled, _ := json.Marshal(data)
	w.Write(marshaled)
}
