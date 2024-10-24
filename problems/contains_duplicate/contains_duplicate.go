package contains_duplicate

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
