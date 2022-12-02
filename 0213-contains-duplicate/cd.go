package containsduplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	flag := false

	for _, item := range nums {
		if ok := m[item]; ok {
			flag = true
		} else {
			m[item] = true
		}
	}

	return flag
}
