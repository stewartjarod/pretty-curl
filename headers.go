package main

import "fmt"

// Header prints the header
func Header(header string) {
	hl := "================================================================================"
	fmt.Println("\x1b[34m")
	fmt.Println(hl)
	fmt.Println(header)
	fmt.Println(hl, "\x1b[0m")
}
