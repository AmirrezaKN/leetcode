package longestsubstringwithoutrepeatingcharacters

import "math"

func lengthOfLongestSubstring(s string) int {
	max := 0
	left := 0
	seen := make(map[rune]bool)

	for _, r := range s {
		for seen[r] {
			delete(seen, rune(s[left]))
			left++
		}
		seen[r] = true
		max = int(math.Max(float64(len(seen)), float64(max)))
	}
	return max
}
