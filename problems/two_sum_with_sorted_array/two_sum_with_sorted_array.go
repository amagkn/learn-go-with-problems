package two_sum

/*
TwoSumWithSortedArray

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order.
Find two numbers such that they add up to a specific target number.
Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
*/
func TwoSumWithSortedArray(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		sum := nums[l] + nums[r]

		if sum == target {
			return []int{l + 1, r + 1}
		}

		if sum > target {
			r--
		}

		if sum < target {
			l++
		}
	}

	return nil
}
