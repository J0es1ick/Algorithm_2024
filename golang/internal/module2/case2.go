package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSortTwoArr(arr, n []int) []int {
	for i := 0; i <= len(n)-1; i++ {
		for j := 0; j < len(n)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func bubbleSortMap(arr []int, n map[int][]int) [] int { 
	for i := 0; i <= len(n)-1; i++ {
		for j := 0; j < len(n)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func Case2() {
	var n int
	fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	products := make(map[int][]int)
	for i := 1; i <= n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")
		str_arr := strings.Split(line, " ")
		valID, err := strconv.Atoi(str_arr[0])
		if err != nil {
			panic(err)
		}
		valPrice, err := strconv.Atoi(str_arr[1])
		if err != nil {
			panic(err)
		}
		_, ok := products[valPrice]
		if !ok {
			products[valPrice] = []int{}
		}
		products[valPrice] = append(products[valPrice], valID)
	}
	keys := make([]int, len(products))
	i := 0
	for k := range products {
    	keys[i] = k
    	i++
	}
	bubbleSortMap(keys, products)
	for _, k := range keys {
		price := k
		bubbleSortTwoArr(products[k], products[k])
		for i := 0; i < len(products[k]); i++ {
			id := products[k][i]
			fmt.Println(id, price)
		}
	}
}