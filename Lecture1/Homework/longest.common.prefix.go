package Homework

// TC: O(n*m) MC: O(1)

func longestCommonPrefix(strs []string) string {
	var ans string
	for i, _ := range strs[0] {
		//fmt.Println(i)
		for _, val2 := range strs {
			if i == len(val2) || val2[i] != strs[0][i] {
				return ans
			}
		}
		ans += string(strs[0][i])
	}
	//fmt.Println(ans)
	return ans

}
