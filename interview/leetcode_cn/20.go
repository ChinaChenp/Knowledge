//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-parentheses

//---------------------------------------------
//初始化栈 S。
//一次处理表达式的每个括号。
//如果遇到开括号，我们只需将其推到栈上即可。这意味着我们将稍后处理它，让我们简单地转到前面的 子表达式。
//如果我们遇到一个闭括号，那么我们检查栈顶的元素。如果栈顶的元素是一个 相同类型的 左括号，那么我们将它从栈中弹出并继续处理。否则，这意味着表达式无效。
//如果到最后我们剩下的栈中仍然有元素，那么这意味着表达式无效。
//

package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

func isValid(str string) bool {
	mappings := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	s := linkedliststack.New()
	for _, v := range str {
		if v == '(' || v == '[' || v == '{' {
			s.Push(v)
		} else if v == ')' || v == ']' || v == '}' {
			top, _ := s.Peek()
			if top.(rune) != mappings[v] {
				return false
			}
			s.Pop()
		}
	}
	return s.Empty()
}

func main() {
	str := "{[[]{}]}()(){"
	fmt.Println(isValid(str))
}
