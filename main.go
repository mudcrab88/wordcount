package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		io.WriteString(os.Stderr, "No arguments!\n")
		os.Exit(1)
	}
	str := strings.TrimSpace(os.Args[1])
	space := 32
	count := 0
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

	fmt.Println(count)
}
