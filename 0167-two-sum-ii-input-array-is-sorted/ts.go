package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	right := len(numbers) - 1
	for left := 0; left < len(numbers) || right >= 0; {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}

		if numbers[left]+numbers[right] > target {
			right--
			continue
		} else {
			left++
			continue
		}
	}

	return nil
}
