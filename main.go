package main

import (
	"brackets/brackets"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		br := brackets.MustNewBracket()
		res := br.Validate(scanner.Text())
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
		return
	}
}
