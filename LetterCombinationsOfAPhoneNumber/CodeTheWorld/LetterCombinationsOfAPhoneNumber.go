package CodeTheWorld

func LetterCombinations(digits string) []string {
	numberCharMap := [8][4]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}

	var res, tmp []string
	for _, value := range digits {
		index := value - '2'
		tmp = []string{}
		for _, val := range numberCharMap[index] {
			if 0 == val {
				continue
			}
			if 0 == len(res) {
				tmp = append(tmp, string(val))
			} else {
				for _, tmpVal := range res {
					tmpVal += string(val)
					tmp = append(tmp, tmpVal)
				}
			}
		}
		res = tmp
	}

	return res
}
