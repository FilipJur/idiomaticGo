package main

import "fmt"

func main() {
	result := div(5, 2)
	fmt.Println(result)
}

func div(num, denom int) int {
	if denom == 0 {
		return 0
	} else {
		return num / denom
	}
}
