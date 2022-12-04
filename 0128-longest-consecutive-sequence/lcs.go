package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	existence := make(map[int]bool)

	for _, val := range nums {
		existence[val] = true
	}

	max := 0
	if len(nums) > 0 {
		max = 1
	}

	for _, num := range nums {
		if existence[num-1] && !existence[num+1] {
			count := 1

			for j := 1; existence[num-j]; j++ {
				count++
			}

			if count > max {
				max = count
			}
		}
	}

	return max
}
