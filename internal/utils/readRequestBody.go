package utils

import (
	"encoding/json"
	"io"
)

func ReadRequestBodyMap[T1 comparable, T2 any](body io.ReadCloser) (map[T1]T2, error) {
	reqBody, err := io.ReadAll(body)

	if err != nil {
		return nil, err
	}

	bodyData := make(map[T1]T2)
	err = json.Unmarshal(reqBody, &bodyData)

	if err != nil {
		return nil, err
	}

	return bodyData, nil
}
