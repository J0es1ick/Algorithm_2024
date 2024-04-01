package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge3(a, b []int) []int {
    final := []int{}
    i, j := 0, 0
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
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
    return final
}

func mergeSort3(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
	left := mergeSort3(arr[:len(arr)/2])
   	right := mergeSort3(arr[len(arr)/2:])
    merged := merge3(left, right)
	return merged
}

func Case5() {
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
	arr = mergeSort3(arr)
	count := 1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			count++
		}
	}
	fmt.Println(count)
}