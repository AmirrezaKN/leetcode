package validparentheses

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.

// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	stack := []string{}

	for _, subStr := range s {
		if string(subStr) == "(" {
			stack = append(stack, string(subStr))
			continue
		}

		if string(subStr) == ")" {
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != "(" {
				return false
			}
			stack = stack[:len(stack)-1]
		}

		if string(subStr) == "{" {
			stack = append(stack, string(subStr))
			continue
		}
		if string(subStr) == "}" {
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		}

		if string(subStr) == "[" {
			stack = append(stack, string(subStr))
			continue
		}
		if string(subStr) == "]" {
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != "[" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
