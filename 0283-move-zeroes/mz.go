package moveZeroes

func moveZeroes(nums []int) {
	left, right := 0, 0

	for ; right < len(nums); right++ {
		if nums[right] != 0 {
			temp := nums[right]
			nums[right] = nums[left]
			nums[left] = temp
			left++
		}
	}
}
