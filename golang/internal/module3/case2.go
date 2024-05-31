package module3

import (
	"fmt"
	"strings"
)

func isCyclicShift(s, t string) int {
	if len(s) != len(t) {
		return -1
	}
	
	doubleT := t + t
	idx := strings.Index(doubleT, s)
	if idx == -1 {
		return -1
	}
	return idx
}

func Case2() {
	var S, T string
	fmt.Scanln(&S)
	fmt.Scanln(&T)

	result := isCyclicShift(S, T)
	fmt.Println(result)
}