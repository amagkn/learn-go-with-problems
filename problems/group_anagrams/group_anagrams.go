package group_anagrams

import (
	"learn-go-with-problems/helpers"
	"sort"
)

func GroupAnagrams(strs []string) [][]string {
	dict := map[string][]string{}

	for _, str := range strs {
		key := helpers.SortString(str)

		if _, ok := dict[key]; ok {
			dict[key] = append(dict[key], str)
		} else {
			dict[key] = []string{str}
		}
	}

	values := helpers.MapValues(dict)

	sort.Slice(values, func(i, j int) bool {
		return len(values[i]) < len(values[j])
	})

	return values
}
