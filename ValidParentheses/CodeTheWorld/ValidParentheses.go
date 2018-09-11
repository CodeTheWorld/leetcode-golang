package CodeTheWorld

func IsValid(s string) bool {
	length := len(s)
	var stack = make([]int, length)
	index := 0
	var value int

	for _, v := range s {
		value = getSymbolValue(v)
		if 0 == value {
			return false
		}
		if index > 0 && 0 == value+stack[index-1] {
			index--
		} else {
			stack[index] = value
			index++
		}
	}

	return 0 == index
}

func getSymbolValue(c rune) int {
	var value int
	switch c {
	case '(':
		value = -1
		break
	case ')':
		value = 1
		break
	case '{':
		value = -2
		break
	case '}':
		value = 2
		break
	case '[':
		value = -3
		break
	case ']':
		value = 3
		break
	default:
		value = 0
		break
	}
	return value
}
