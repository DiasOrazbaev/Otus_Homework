package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	if len(str) == 0 {
		return []string{}
	}
	words := strings.Fields(str)

	dic := make(map[string]int)
	keys := make([]string, 0)

	for _, word := range words {
		dic[word]++
	}

	for key := range dic {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if dic[keys[i]] == dic[keys[j]] {
			return keys[i] < keys[j]
		}
		return dic[keys[i]] > dic[keys[j]]
	})

	if len(keys) > 10 {
		keys = keys[:10]
	}

	return keys
}
