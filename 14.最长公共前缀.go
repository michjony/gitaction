package main

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (32.55%)
 * Total Accepted:    68.1K
 * Total Submissions: 208.4K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */
func longestCommonPrefix(strs []string) string {

	if len(strs) < 1 {
		return ""
	}
	startStr := strs[0]
	var commonStr string = startStr
	for index := 1; index < len(strs); index++ {
		comparestr := strs[index]

		for i := 0; i < len(commonStr); i++ {
			if len(comparestr) > i && commonStr[i] == comparestr[i] {

			} else {
				commonStr = commonStr[0:i]
			}

		}

	}
	return commonStr

}
