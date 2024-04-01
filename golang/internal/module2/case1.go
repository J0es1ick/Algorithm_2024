package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Case1() {
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
	num_swaps := 0
	for i := 0; i <= n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				num_swaps += 1
				fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
			}
		}
	}
	if num_swaps == 0 {
		fmt.Println("0")
	}
}
