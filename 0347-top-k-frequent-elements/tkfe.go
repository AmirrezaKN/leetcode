package topkfrequentelements

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num] += 1
	}

	res := [10000][]int{}
	for num, count := range counts {
		res[count] = append(res[count], num)
	}

	output := []int{}
	amount := 0
	for i := 10000 - 1; i >= 0; i-- {
		if amount == k {
			break
		}
		if len(res[i]) > 0 {
			amount += len(res[i])
			output = append(output, res[i]...)
		}
	}

	return output
}
