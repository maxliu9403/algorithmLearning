package main

import (
	"fmt"
	"strconv"
)


// 转换为后缀表达
func evalRPN(tokens []string) int {
	number := []int{}
	for _, val := range tokens{
		l := len(number)
		switch val {
		case "+":
			number  = append(number[:l -2], number[l-2] + number[l-1])
		case "-":
			number  = append(number[:l -2], number[l-2] - number[l-1])
		case "*":
			number  = append(number[:l -2], number[l-2] * number[l-1])
		case "/":
			number  = append(number[:l -2], number[l-2] / number[l-1])
		default:
			num, _ := strconv.Atoi(val)
			number  = append(number, num)
		}
	}
	return number[0]
}

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	num := evalRPN(tokens)
	fmt.Println(num)
}
