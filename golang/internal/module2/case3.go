package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge(a, b []int) []int {
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

func mergeSort(arr []int, idx int) []int {
    if len(arr) < 2 {
        return arr
    }
    mid := len(arr) / 2
	left := mergeSort(arr[:mid], idx)
   	right := mergeSort(arr[mid:], idx+mid)
    merged := merge(left, right)
	fmt.Println(idx, idx + len(arr) - 1, merged[0], merged[len(merged)-1])
	return merged
}

func Case3() {
	var n int
	fmt.Scanln(&n)
	idx := 1
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
	fmt.Println(strings.Trim(fmt.Sprint(mergeSort(arr, idx)), "[]"))
}
