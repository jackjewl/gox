package others

import "gox/utils"

//递归法求最大公共字符串长度
func Lcs1(a []rune, b []rune) int {

	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	if a[len(a)-1] == b[len(b)-1] {
		return Lcs1(a[:len(a)-1], b[:len(b)-1]) + 1
	} else {
		return utils.Max[int](Lcs1(a[:len(a)-1], b), Lcs1(a, b[:len(b)-1]))
	}

}

//动态规划求最大公共子字符串长度
func Lcs2(a, b []rune) int {

	table := make([][]int, len(a)+1)
	for index, _ := range table {
		table[index] = make([]int, len(b)+1)
	}

	for i := 1; i < len(table); i++ {
		for j := 1; j < len(table[i]); j++ {
			if a[i-1] == b[j-1] {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				table[i][j] = utils.Max[int](table[i-1][j], table[i][j-1])
			}

		}
	}
	return table[len(a)][len(b)]
}
