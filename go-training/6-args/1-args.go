package main

import (
	"fmt"
	"os"
)

func main() {

	os.Args[0] = "sfs"
	a := os.Args[0:]
	fmt.Printf("%T", os.Args)
	fmt.Println()
	fmt.Println(len(os.Args))
	fmt.Println(a)
}
