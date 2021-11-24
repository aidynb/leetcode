package lc

type stack struct {
	stack  []rune
	length int
}

func (s *stack) peek() rune {
	if s.length > 0 {
		return s.stack[s.length-1]
	}
	return 0
}

func (s *stack) pop() rune {
	lastElement := s.stack[s.length-1]
	s.stack = s.stack[:s.length-1]
	s.length--
	return lastElement
}

func (s *stack) push(element rune) {
	s.stack = append(s.stack, element)
	s.length++
}

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	stack := stack{}

	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, bracket := range s {
		value, exists := bracketMap[bracket]
		if exists {
			if stack.peek() == value {
				stack.pop()
			} else {
				return false
			}
		} else {
			stack.push(bracket)
		}
	}

	return stack.length == 0
}
