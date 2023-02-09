package main

import "fmt"

func generateSequences(current string, remaining int) {
	if remaining == 0 {
		fmt.Println(current)
		return
	}

	generateSequences(current+"+", remaining-1)
	generateSequences(current+"-", remaining-1)
	generateSequences(current+"*", remaining-1)
}

func main() {
	generateSequences("", 8)
}
