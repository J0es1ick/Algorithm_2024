package module4

import (
	"fmt"
)

func minRemoveToMakeValid(s string) int {
	stack := []rune{}
	count := 0

	for _, char := range s {
		if char == '(' {
			stack = append(stack, char)
		} else if char == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				count++
			}
		}
	}

	return count + len(stack)
}

func Case1() {
	var s string
	fmt.Scan(&s)

	result := minRemoveToMakeValid(s)
	fmt.Println(result)
}