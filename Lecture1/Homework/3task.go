package Homework

import "sort"

// TC: O(n*logn) MC(O(1))

func compareWithSort(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for idx, _ := range a {
		if a[idx] != b[idx] {
			return false
		}
	}

	return true
}
