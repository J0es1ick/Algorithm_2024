package module3

import (
	"fmt"
	"hash/crc32"
)

func minRepeatLength(s string) int {
	for i := 1; i <= len(s); i++ {
		substr := s[:i]
		repeats := len(s) / i
		newStr := ""
		for j := 0; j < repeats; j++ {
			newStr += substr
		}
		if crc32.ChecksumIEEE([]byte(newStr)) == crc32.ChecksumIEEE([]byte(s)) {
			return i
		}
	}
	return len(s)
}

func Case4() {
	var s string
	fmt.Scanln(&s)

	result := minRepeatLength(s)
	fmt.Println(result)
}