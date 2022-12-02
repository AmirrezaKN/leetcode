package groupanagrams

func groupAnagrams(strs []string) [][]string {
	result := make(map[[26]int][]string)

	for _, str := range strs {
		count := [26]int{}

		for _, r := range str {
			var aRune rune = rune('a')
			count[r-aRune] += 1
		}

		if _, ok := result[count]; ok {
			result[count] = append(result[count], str)
		} else {
			result[count] = []string{str}
		}
	}

	output := make([][]string, 0)

	for _, val := range result {
		output = append(output, val)
	}

	return output
}
