package merge_intervals

import "sort"

/*
MergeIntervals

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/
func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		lastMerged := merged[len(merged)-1]
		current := intervals[i]

		if lastMerged[1] >= current[0] {
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}
