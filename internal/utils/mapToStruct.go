package utils

import (
	"encoding/json"
)

func MapToStruct(mapped map[string]any, Struct any) {
	jsonString, _ := json.Marshal(mapped)
	_ = json.Unmarshal(jsonString, &Struct)
}

func CastAndCompare[T any](rawMessage any) (T, bool) {
	castResult, ok := rawMessage.(map[string]any)

	var retVal T
	if !ok || !CompareJSONToStruct(castResult, retVal) {
		return retVal, false
	}

	MapToStruct(castResult, &retVal)

	return retVal, true
}
