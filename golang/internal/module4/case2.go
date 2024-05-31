package module4

import (
	"fmt"
	"sort"
)

func Case2() {
	var N int
	fmt.Scan(&N)

	var pathA []int
	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		pathA = append(pathA, num)
	}

	var pathB []int
	for _, num := range pathA {
		pathB = append(pathB, num)
	}

	sort.Ints(pathB)

	flag := true
	for i := 0; i < N; i++ {
		if pathB[i] != i+1 {
			flag = false
			break
		}
	}

	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}