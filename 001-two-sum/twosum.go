package twosum

func twoSum(nums []int, target int) []int {
	diff := make(map[int]int)

	for i, num := range nums {
		if j, ok := diff[num]; ok && i != j {
			return []int{i, j}
		}
		diff[target-num] = i
	}

	return nil
}
