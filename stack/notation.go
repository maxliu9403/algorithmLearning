package main

import (
	"fmt"
	"strconv"
)

// 实现stack
type Constructor struct {
	stack []interface{}
	len   int
}

func Stack() *Constructor {
	s := &Constructor{}
	s.stack = make([]interface{}, 0)
	s.len = 0
	return s
}

func (s *Constructor) Len() int {
	return s.len
}

func (s *Constructor) IsEmpty() bool {
	return s.len == 0
}

func (s *Constructor) Pop() (element interface{}) {
	element = s.stack[s.len-1]
	s.stack = s.stack[:s.len-1]
	s.len--
	return
}

// 压栈
func (s *Constructor) Push(element interface{}) {
	s.stack = append(s.stack, element)
	s.len++
}

// 栈顶元素
func (s *Constructor) Top() interface{} {
	return s.stack[s.len-1]
}



// 运算符号优先级
var priority = map[string]int{
	"(": 1, ")": 1,
	"*": 2, "/": 2,
	"+": 3, "-": 3,
}

func isNum(s string) bool {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return false
	}
	return true
}

// 运算符号判断
func isOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

// 括号判断
func isBracket(s string) bool {
	if s == "(" || s == ")" {
		return true
	}
	return false
}


// 中缀转后缀
func IN2RPN(infix []string) (reversePolish []string) {

	// 运算符栈S1和储存中间结果栈S2；
	s1, s2 := Stack(), Stack()

	// infix := []string{"4.99", "*", "5.99", "+", "6.99", "*", "1.06"}

	for i := 0; i < len(infix); i++ {
		s := string(infix[i])

		// 如果是数字压入s2
		if isNum(s) {
			s2.Push(s)

		// 判断是否是运算符
		} else if isOperator(s) {

			// 如果s1为空，或栈顶运算符为左括号“(”，则直接将此运算符入栈
			if s1.IsEmpty() || s1.Top().(string) == "(" ||
				priority[s] < priority[s1.Top().(string)] {
				s1.Push(s)
			} else {
				s2.Push(s1.Pop())
				i--
			}
		} else if isBracket(s) {
			// 如果是左括号“(”，则直接压入S1；
			if s == "(" {
				s1.Push(s)

				// 如果是右括号“)”，则依次弹出S1栈顶的运算符，并压入S2，直到遇到左括号为止，此时将这一对括号丢弃
			} else {
				for !s1.IsEmpty() && s1.Top().(string) != "(" {
					s2.Push(s1.Pop())
				}
				if s1.Top().(string) == "(" {
					s1.Pop()
				}
			}
		} else {
			panic(fmt.Sprintf("error input: %s", s))
		}
	}

	// 将S1中剩余的运算符依次弹出并压入S2
	for !s1.IsEmpty() {
		s2.Push(s1.Pop())
	}

	// 依次弹出S2中的元素并输出，结果的逆序即为中缀表达式对应的后缀表达式
	res, index := make([]string, s2.Len(), s2.Len()), s2.Len()-1
	for !s2.IsEmpty() {
		res[index] = s2.Pop().(string)
		index--
	}
	return res
}

func main() {
	infix := []string{"1", "*", "2", "+", "3", "*", "4"}
	RPN := IN2RPN(infix)
	fmt.Println(RPN)

}