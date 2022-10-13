package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	count := 0
	if len(os.Args) == 2 {
		str := strings.TrimSpace(os.Args[1])
		space := 32

		prev := int(str[0])

		if len(str) > 0 {
			count = 1
			for _, val := range str {
				if int(val) == space && prev != space {
					count += 1
				}
				prev = int(val)
			}
		}
	}

	fmt.Println(count)
}
