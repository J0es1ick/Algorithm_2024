package module3

import (
	"fmt"
)

func findSubstrIdxs(s, t string) []int {
	idxs := []int{}
	for i := 0; i <= len(s)-len(t); i++ {
		if s[i:i+len(t)] == t {
			idxs = append(idxs, i)
		}
	}
	return idxs
}

func Case1() {
	var S, T string
	fmt.Scanln(&S)
	fmt.Scanln(&T)

	idxs := findSubstrIdxs(S, T)

	for _, idx := range idxs {
		fmt.Print(idx, " ")
	}
}