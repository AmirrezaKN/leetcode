package longestcommonprefix

import "strings"

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lower-case English letters.

func longestCommonPrefix(strs []string) string {
	result := strs[0]

	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], result) {
			if result == "" {
				return ""
			}
			result = result[:len(result)-1]
		}
	}

	return result
}
