package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	if name := strings.Join(os.Args[1:], " "); name != "" {
		print("Hello, " + name)
		return
	}

	print("Hello, Anonymous")
}

func print(content string) {
	fmt.Println(content)
}
