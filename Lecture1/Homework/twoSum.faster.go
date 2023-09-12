package Homework

//TC: O(n) MC: O(1)

func twoSumFaster(nums []int, target int) []int {
	var (
		hashMap = map[int]int{}
		dif     int
	)

	for idx, val := range nums {
		//formula to calculate the difference
		dif = target - val

		idx2, ok := hashMap[dif]
		//if the second number exists in hashMap then it's right
		if ok {
			return []int{idx, idx2}
		}
		hashMap[val] = idx
	}
	return []int{}
}
