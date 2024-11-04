package group_anagrams

import (
	"learn-go-with-problems/my_helpers"
	"sort"
)

/*
GroupAnagrams

Given an array of strings strs, group the
anagrams together. You can return the answer in any order.
*/
func GroupAnagrams(strs []string) [][]string {
	dict := map[string][]string{}

	for _, str := range strs {
		key := my_helpers.SortString(str)

		if _, ok := dict[key]; ok {
			dict[key] = append(dict[key], str)
		} else {
			dict[key] = []string{str}
		}
	}

	values := my_helpers.MapValues(dict)

	sort.Slice(values, func(i, j int) bool {
		return len(values[i]) < len(values[j])
	})

	return values
}
