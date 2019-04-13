package main

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.32%)
 * Total Accepted:    63.5K
 * Total Submissions: 170.1K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */
func isValid(s string) bool {
	stack := []string{}
	for _, va := range s {
		v := string(va)
		if len(stack) == 0 {
			stack = append(stack, v)
		} else {
			le := len(stack)
			// get element
			str := stack[le-1]

			if (str == "{" && v == "}") || (str == "}" && v == "{") || (str == "[" && v == "]") || (str == "]" && v == "[") || (str == "(" && v == ")") || (str == ")" && v == "(") {

				stack = stack[0 : le-1]
			} else {

				stack = append(stack, v)
			}

		}

	}
	if len(stack) > 0 {
		return false
	}
	return true
}
