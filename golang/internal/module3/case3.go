package module3

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"log"
	"os"
)

func Case3() {
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        log.Fatal(scanner.Err())
    }
    s := scanner.Text()

    bestK := 0

    for x := 0; x < len(s); x++ {
            for i := x+1; i < len(s); i++ {
            temp := 1
            arr := []byte(s[x:i])
            hash := crc32.ChecksumIEEE(arr)
            for j := i; j <= len(s)-(i-x); j += (i-x) {
                if crc32.ChecksumIEEE([]byte(s[j:j+(i-x)])) != hash {
                    break
                }
                temp++
            }
            
            if temp > bestK {
                bestK = temp
            }
        }
    }

    fmt.Println(bestK)
}