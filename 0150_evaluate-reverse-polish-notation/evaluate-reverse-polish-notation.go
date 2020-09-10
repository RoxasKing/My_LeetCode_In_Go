package main

import "strconv"

/*
  根据 逆波兰表示法，求表达式的值。
  有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

  说明：
    整数除法只保留整数部分。
    给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

  逆波兰表达式：
  逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。
    平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
    该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。
  逆波兰表达式主要有以下两个优点：
    去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
    适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/evaluate-reverse-polish-notation
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack
func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		if token != "+" && token != "-" && token != "*" && token != "/" {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
			continue
		}
		last := len(stack) - 1
		num1, num2 := stack[last-1], stack[last]
		switch token {
		case "+":
			stack = append(stack, num1+num2)
		case "-":
			stack = append(stack, num1-num2)
		case "*":
			stack = append(stack, num1*num2)
		case "/":
			stack = append(stack, num1/num2)
		}
	}
	return stack[0]
}
