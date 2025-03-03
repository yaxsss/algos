package stack

func isValid(s string) bool {
	stack := []byte{}
	pop := func() {
		stack = stack[0 : len(stack)-1]
	}
	push := func(c byte) {
		stack = append(stack, c)
	}
	peek := func() byte {
		if len(stack) == 0 {
			return 0
		}
		return stack[len(stack)-1]
	}
	for c := range s {
		switch s[c] {
		case '(', '[', '{':
			push(s[c])

		case ')', ']', '}':
			pair := byte('(')
			if s[c] == ']' {
				pair = '['
			} else if s[c] == '}' {
				pair = '{'
			}
			if peek() != pair {
				return false
			}
			pop()
		}
	}
	return len(stack) == 0
}
