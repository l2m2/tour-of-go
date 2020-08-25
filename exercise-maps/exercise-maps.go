// 练习：映射
// 实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。
// 函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
	// WordCount("bar foo bar ")
}
