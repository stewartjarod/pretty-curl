package main

import "fmt"

// PPKeyVal pretifies a key / value pair
func PPKeyVal(s string, v string) {
	fmt.Printf("\x1b[30;1m%v \x1b[30;0m%s\n", s, v)
	return
}

// PPrintCode pretifies response codes
func PPrintCode(title string, code int) {
	if code == 200 || code == 201 || code == 202 {
		fmt.Printf("\x1b[30;1m%s \x1b[22;32m%d\x1b[30;0m\n", title, code)
	} else {
		fmt.Printf("\x1b[30;1m%s \x1b[30;0m%d\x1b[30;0m\n", title, code)
	}
	return
}
