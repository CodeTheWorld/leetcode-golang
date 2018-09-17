package CodeTheWorld

var res []string

func GenerateParenthesis(n int) []string {
	res = []string{}
	if 0 >= n {
		return res
	}
	dfs(0, 0, n, "")
	return res
}

func dfs(left int, right int, n int, str string) {
	if n == left && n == right {
		res = append(res, str)
		return
	}

	if left < n {
		dfs(left+1, right, n, str+"(")
	}
	if left > right {
		dfs(left, right+1, n, str+")")
	}
}
