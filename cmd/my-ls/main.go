package main

import (
	"fmt"
	internal "my-ls/internal/ls"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {

	} else {
		path := os.Args[1]
	}

	files := internal.RetrieveFile(path)
	fmt.Println(files)
}
