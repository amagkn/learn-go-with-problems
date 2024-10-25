package top_k_frequent_elements

import (
	"learn-go-with-problems/my_helpers"
	"sort"
)

/*
TopKFrequent

Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
*/
func TopKFrequent(nums []int, k int) []int {
	dict := map[int]int{}

	for _, num := range nums {
		dict[num]++
	}

	keys := my_helpers.MapKeys(dict)

	sort.Slice(keys, func(i, j int) bool {
		return dict[keys[i]] > dict[keys[j]]
	})

	return keys[:k]
}
