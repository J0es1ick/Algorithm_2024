package module2

import (
	"fmt"
	"strconv"
)

func Case7() {
	var n int
	fmt.Scanln(&n) 

	str := make([]string, n) 
	buckets := make([][]string, 10)

	for i := 0; i < n; i++ {
		fmt.Scan(&str[i])
	}

	fmt.Println("Initial array:")
	for i := 0; i < len(str); i++ {
		if i < len(str)-1 {
			fmt.Print(str[i], ", ")
		} else {
			fmt.Println(str[i])
		}
	}

	maxLen := 0
	for _, s := range str {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	} 

	j := 1
	for i := maxLen - 1; i >= 0; i-- {
		fmt.Println("**********")
		fmt.Println("Phase", j)
		j++

		for _, s := range str {
			digit := 0
			if i < len(s) {
				digit, _ = strconv.Atoi(string(s[i]))
			}
			buckets[digit] = append(buckets[digit], s)
		}

		for j, bucket := range buckets {
			fmt.Printf("Bucket %d: ", j)
			if len(bucket) > 0 {
				for i := 0; i < len(bucket); i++ {
					if i < len(bucket)-1 {
						fmt.Print(bucket[i], ", ")
					} else {
						fmt.Println(bucket[i])
					}
				}
			} else {
				fmt.Println("empty")
			}
		}
		
		idx := 0
		for j := 0; j < len(buckets); j++ {
			for _, s := range buckets[j] {
				str[idx] = s
				idx++
			}
			buckets[j] = nil
		}
	}

	fmt.Println("**********")
	fmt.Println("Sorted array:")
	for i := 0; i < len(str); i++ {
		if i < len(str)-1 {
			fmt.Print(str[i], ", ")
		} else {
			fmt.Println(str[i])
		}
	}
}