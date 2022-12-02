package validanagram

func isAnagram(s string, t string) bool {
	firstValues := make(map[rune]int)
	secondValues := make(map[rune]int)
	flag := true

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

	for rune, value := range secondValues {
		if value != firstValues[rune] {
			flag = false
			break
		}
	}

	return flag
}
