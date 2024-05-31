package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var arr []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		arr = append(arr, num)
	}

	for i := 0; i < n; i++ {
		var result int
		left, right := i+1, n-1
		for left <= right {
			mid := (left + right) / 2
			if arr[mid] < arr[i] {
				arr[i] = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if result == n-1 {
			arr[n-1] = -1
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}