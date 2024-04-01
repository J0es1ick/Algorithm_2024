package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertString(arr []int, count int) ([]int, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	str_arr := strings.Split(line, " ")
	for idx, val := range str_arr {
		arr[idx], err = strconv.Atoi(val)
	}
	return arr, err
}

func Case6() {
	var countOfKinds int
	fmt.Scanln(&countOfKinds)
	kinds := make([]int, countOfKinds)
	kinds, err := convertString(kinds, countOfKinds)
	if err != nil {
		panic(err)
	}
	
	var countOfOrders int
	fmt.Scanln(&countOfOrders)
	orders := make([]int, countOfOrders)
	orders, err = convertString(orders, countOfOrders)
	if err != nil {
		panic(err)
	}
	
	sum := make([]int, countOfKinds)
	for i := 0; i < countOfOrders; i++ {
		order := orders[i] - 1
		sum[order] += 1
	}
	
	for j := 0; j < countOfKinds; j++ {
		if kinds[j] >= sum[j] {
			fmt.Println("no")
		}else{
			fmt.Println("yes")
		}
	}
}