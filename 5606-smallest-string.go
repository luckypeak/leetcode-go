package leetcode_go

import (
	"bytes"
)

/**
 贪心算法,尽可能多的z，同时尽可能的a保证字典序最小，在选择完z,a 之后剩余一位选择其他元素补齐
 注意大量字符串拼接，直接使用+ 拼接耗时比较长
 字符串拼接方法 1, +, 2.strings.Join, 3. bytes.Buffer
 */
func getSmallestString(n int, k int) string {
	e := 0
	var b bytes.Buffer
	for k - e * 26 - (n-e-1) > 26 { // 减1是因为需要在a到z之间留一位元素
		e ++
	}
	remain := n-e-1
	for i := 0; i < remain; i++{
		b.WriteString("a")
	}
	b.WriteRune(rune(k-e*26-remain-1+'a'))

	for i := 0; i < e; i++{
		b.WriteString("z")
	}

	return  b.String()

}