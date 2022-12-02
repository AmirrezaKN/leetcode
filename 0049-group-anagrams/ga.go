package groupanagrams

func groupAnagrams(strs []string) [][]string {
	result := make(map[[26]int][]string)

	for _, str := range strs {
		count := [26]int{}

		for _, r := range str {
			count[r-rune('a')] += 1
		}
		result[count] = append(result[count], str)
	}

	output := make([][]string, 0)

	for _, val := range result {
		output = append(output, val)
	}

	return output
}
