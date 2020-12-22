package main

/*
  验证给定的字符串是否可以解释为十进制数字。

  例如:
    "0" => true
    " 0.1 " => true
    "abc" => false
    "1 a" => false
    "2e10" => true
    " -90e3   " => true
    " 1e" => false
    "e3" => false
    " 6e-1" => true
    " 99e2.5 " => false
    "53.5e93" => true
    " --6 " => false
    "-+3" => false
    "95a54e53" => false

  说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：
    数字 0-9
    指数 - "e"
    正/负号 - "+"/"-"
    小数点 - "."
  当然，在输入中，这些字符的上下文也很重要。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/valid-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isNumber(s string) bool {
	for s != "" && s[0] == ' ' {
		s = s[1:]
	}

	if s != "" && (s[0] == '-' || s[0] == '+') {
		s = s[1:]
	}

	var containNumber bool
	for s != "" && '0' <= s[0] && s[0] <= '9' {
		containNumber = true
		s = s[1:]
	}

	if s != "" && s[0] == '.' {
		s = s[1:]
	}

	for s != "" && '0' <= s[0] && s[0] <= '9' {
		containNumber = true
		s = s[1:]
	}

	if s != "" && containNumber && s[0] == 'e' {
		s = s[1:]
		containNumber = false
		if s != "" && (s[0] == '-' || s[0] == '+') {
			s = s[1:]
		}
		for s != "" && '0' <= s[0] && s[0] <= '9' {
			containNumber = true
			s = s[1:]
		}
	}

	for s != "" && s[0] == ' ' {
		s = s[1:]
	}

	return s == "" && containNumber
}