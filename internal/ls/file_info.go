//This file will contain the logic for retrieving file and directory metadata.
// It uses Go’s os and syscall packages (since the os/exec package is not allowed).
//Functions here will retrieve file sizes, permissions, timestamps, etc.

package internal

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func RetrieveFile(path string) []string {

// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(file)
// }
