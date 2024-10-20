package main

import (
	"fmt"
	"log"
	internal "my-ls/internal/ls"
	"os"
)

func main() {
	args := os.Args[1:] // Retrieve arguments from command line

	fmt.Printf("%q\n", args)
	// Extract flags and paths from user arguments
	// Handle errors, if encountered
	flag, path, err := internal.SortArgs(args)
	if err != nil {
		log.Fatal(err)
	}

	files := internal.RetrieveFileInfo(path)
	fmt.Println(files)
	fmt.Println(flag)
}
