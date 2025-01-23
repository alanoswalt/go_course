//Print_file

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	f, err := os.Open(os.Args[1])

	if err != nil {
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
