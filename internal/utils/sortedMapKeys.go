package utils

import "sort"

func SortedMapKeys[T any](mapped map[string]T) []string { // дженерик тут нужен, ибо ошибит с мизматчем типов при map[xxx]any
	keys := make([]string, 0, len(mapped))
	for k := range mapped {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
