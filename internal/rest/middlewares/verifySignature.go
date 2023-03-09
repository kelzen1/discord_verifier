package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/yoonaowo/discord_verifier/internal/models/rest"
	"github.com/yoonaowo/discord_verifier/internal/utils"
	"io"
	"net/http"
)

func CheckSignature(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		answer := &restModels.ErrorAnswer{Error: "unknown"}
		var success bool

		defer func() {
			if !success {
				marshaled, _ := json.Marshal(answer)
				w.WriteHeader(400)
				_, _ = w.Write(marshaled)
			}
		}()

		reqBody, _ := io.ReadAll(r.Body)
		r.Body.Close() //  must close
		r.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		bodyData := make(map[string]string)
		err := json.Unmarshal(reqBody, &bodyData)

		if err != nil {
			answer.Error = err.Error()
			return
		}

		signature, found := bodyData["signature"]

		if !found {
			answer.Error = "signature_not_found"
			return
		}

		delete(bodyData, "signature")

		mapKeys := utils.SortedMapKeys[string](bodyData)
		var hashString string

		for _, key := range mapKeys {
			hashString += key + bodyData[key]
		}

		hashString = utils.HashMD5(hashString + utils.FlagSecretPassword)

		if hashString != signature {
			answer.Error = "wrong_signature"
			return
		}

		success = true
		next.ServeHTTP(w, r)
	})
}
