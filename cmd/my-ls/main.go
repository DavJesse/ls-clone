package main

import (
	"fmt"
	"log"
	internal "my-ls/internal/ls"
	"os"
)

func main() {
	args := os.Args[1:] // Retrieve arguments from command line

	// Extract flags and paths from user arguments
	// Handle errors, if encountered
	flag, path, err := internal.SortArgs(args)
	if err != nil {
		log.Fatal(err)
	}

	files := internal.RetrieveFileInfo(path, false)
	for i := range files {
		fmt.Print("Index: ")
		fmt.Println(files[i].Index)
		fmt.Println()
		fmt.Println()

		fmt.Print("File/Dir Name: ")
		fmt.Println(files[i].DocName)
		fmt.Println()
		fmt.Println()

		fmt.Print("File/Dir Detail: ")
		fmt.Println(files[i].DocPerm)
		fmt.Println()
		fmt.Println()
	}
	//fmt.Println(files)
	fmt.Println(flag)
}
