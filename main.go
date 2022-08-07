package main

import (
	"brackets/brackets"
	"fmt"
)

func main() {
	var bracesStr string
	fmt.Scan(&bracesStr)
	br := brackets.MustNewBracket()
	res := br.Validate(bracesStr)
	if res > 0 {
		fmt.Println(res)
	} else {
		fmt.Println("Success")
	}
	return
}
