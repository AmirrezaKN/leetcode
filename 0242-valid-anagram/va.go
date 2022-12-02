package validanagram

func isAnagram(s string, t string) bool {
	firstValues := make(map[rune]int)
	secondValues := make(map[rune]int)
	flag := true

	if len(t) != len(s) {
		return false
	}

	for _, rune := range s {
		firstValues[rune] += 1
	}

	for _, rune := range t {
		secondValues[rune] += 1
	}

	for rune, value := range firstValues {
		if value != secondValues[rune] {
			flag = false
			break
		}
	}

	return flag
}
