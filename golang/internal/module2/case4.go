package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge2(a, b []int) ([]int, int) {
    final := []int{}
    i, j, count := 0, 0, 0
    for i < len(a) && j < len(b) {
        if a[i] <= b[j] {
            final = append(final, a[i])
            i++
        } else {
            count += len(a) - i
			final = append(final, b[j])
            j++
        }
    }
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }
    return final, count
}

func mergeSort2(arr []int) ([]int, int) {
    if len(arr) < 2 {
        return arr, 0
    }
    mid := len(arr) / 2
	left, a := mergeSort2(arr[:mid])
   	right, b := mergeSort2(arr[mid:])
    merged, c := merge2(left, right)
	return merged, a + b + c
}

func Case4() {
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	if n > 0 {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")
		str_arr := strings.Split(line, " ")
		for idx, val := range str_arr {
			arr[idx], err = strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
		}
	}
	arr, count := mergeSort2(arr)
	fmt.Println(count)
}