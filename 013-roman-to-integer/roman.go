package romantointeger

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.

// 1 <= s.length <= 15
// s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
// It is guaranteed that s is a valid roman numeral in the range [1, 3999].

func romanToInt(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "I" {
			if i+1 < len(s) {
				if string(s[i+1]) == "V" {
					result += 4
					i++
					continue
				}
				if string(s[i+1]) == "X" {
					result += 9
					i++
					continue
				}
			}
			result += 1
			continue
		}

		if string(s[i]) == "X" {
			if i+1 < len(s) {
				if string(s[i+1]) == "L" {
					result += 40
					i++
					continue
				}
				if string(s[i+1]) == "C" {
					result += 90
					i++
					continue
				}
			}

			result += 10
			continue
		}

		if string(s[i]) == "C" {
			if i+1 < len(s) {
				if string(s[i+1]) == "D" {
					result += 400
					i++
					continue
				}
				if string(s[i+1]) == "M" {
					result += 900
					i++
					continue
				}
			}

			result += 100
			continue
		}

		if string(s[i]) == "V" {
			result += 5
		}

		if string(s[i]) == "L" {
			result += 50
		}

		if string(s[i]) == "D" {
			result += 500
		}

		if string(s[i]) == "M" {
			result += 1000
		}

	}

	return result
}
