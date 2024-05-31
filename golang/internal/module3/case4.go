package module3

import (
	"fmt"
)

func minRepeatLength(s string) int {
	for i := 1; i <= len(s); i++ {
		if len(s)%i == 0 {
			substr := s[:i]
			repeats := len(s) / i
			newStr := ""
			for j := 0; j < repeats; j++ {
				newStr += substr
			}
			if newStr == s {
				return i
			}
		}
	}
	return len(s)
}

func Case4() {
	var s string
	fmt.Scanln(&s)

	result := minRepeatLength(s)
	fmt.Println(result)
}