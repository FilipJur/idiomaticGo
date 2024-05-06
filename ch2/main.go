package main

import "fmt"

func exercise1() {
	var i int = 20
	var f float64 = float64(i)

	fmt.Println(i, f)
	fmt.Printf("Type of f: %T\n", f)
}

func exercise2() {
	const value = 1
	var i int = value
	var f float32 = value

	fmt.Println(i, f)
	fmt.Printf("Type of vars: %T, %T\n", i, f)
}

func exercise3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1
	fmt.Println("b:", b, "\nsmallI:", smallI, "\nbigI:", bigI)
}

func main() {
	exercise1()
	exercise2()
	exercise3()
}
