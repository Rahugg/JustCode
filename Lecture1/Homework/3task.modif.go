package Homework

// TC: O(N) MC: O(N)

func compareWithoutSort(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	var a_map, b_map = map[int]int{}, map[int]int{}

	for idx, val := range a {
		a_map[val]++
		b_map[b[idx]]++
	}

	for key, _ := range a_map {
		if a_map[key] != b_map[key] {
			return false
		}
	}

	return true
}
