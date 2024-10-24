package contains_duplicate

/*
ContainsDuplicate

Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.
*/
func ContainsDuplicate(nums []int) bool {
	dict := map[int]bool{}

	for _, n := range nums {
		if _, ok := dict[n]; ok {
			return true
		}

		dict[n] = true
	}

	return false
}
