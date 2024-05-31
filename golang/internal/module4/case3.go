package module4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nearestSmallerOne( arr []int, N int) []int {
	for i := 0; i < N; i++ {
		if i < N-1 {
			for j := i + 1; j < N; j++ {
				if arr[j] < arr[i] {
					arr[i] = j 
					break
				} else if j == N-1 {
					arr[i] = -1
				}
			}
		} else {
			arr[i] = -1
		}
	}
	return arr
}

func Case3() {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	str_arr := strings.Split(line, " ")
	for idx, val := range str_arr {
		arr[idx], _ = strconv.Atoi(val)
	}
	
	fmt.Println(strings.Trim(fmt.Sprint(nearestSmallerOne(arr, N)), "[]"))
}