package main

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现strStr()
 *
 * https://leetcode-cn.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (38.01%)
 * Total Accepted:    46.9K
 * Total Submissions: 123.4K
 * Testcase Example:  '"hello"\n"ll"'
 *
 * 实现 strStr() 函数。
 *
 * 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置
 * (从0开始)。如果不存在，则返回  -1。
 *
 * 示例 1:
 *
 * 输入: haystack = "hello", needle = "ll"
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: haystack = "aaaaa", needle = "bba"
 * 输出: -1
 *
 *
 * 说明:
 *
 * 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
 *
 * 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
 *
 */
func strStr(haystack string, needle string) int {
	if "" == needle {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	start := -1
	last := 0
	flag := false
	for index := 0; index < len(haystack); index++ {
		if last < len(needle) && haystack[index] == needle[last] {
			if !flag {
				flag = true
				start = index
			}
			last++
			if last == len(needle) {
				break
			}
		} else {
			flag = false
			if start >= 0 {
				index = start
			}
			last = 0
			start = -1

		}
	}
	if last < len(needle) {
		return -1
	}
	return start
}
