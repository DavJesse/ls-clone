package main

import (
	"fmt"
	internal "my-ls/internal/ls"
	"os"
)

func main() {
	var path string
	args := os.Args[1:]

	if len(args) == 0 {
		path = "."
	} else {
		
		path = os.Args[1]
	}

	files := internal.RetrieveFile(path)
	fmt.Println(files)
}
