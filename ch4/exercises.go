package main

import (
	"fmt"
	"math/rand"
)

func exercisesCh4() {
	// exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	randomNumbers := []int{}
	for i := 0; i < 100; i++ {
		randomNumber := rand.Intn(100)
		randomNumbers = append(randomNumbers, randomNumber)
	}
	fmt.Println(randomNumbers)
}

func exercise2() {
	randomNumbers := []int{}
	for i := 0; i < 100; i++ {
		randomNumber := rand.Intn(100)
		randomNumbers = append(randomNumbers, randomNumber)
	}
	fmt.Println(randomNumbers)
	for _, number := range randomNumbers {
		switch {
		case number%2 == 0 && number%3 == 0:
			fmt.Println("Six!")
		case number%2 == 0:
			fmt.Println("Two!")
		case number%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("nevermind....")
		}
		// if v%2 == 0 {
		// 	fmt.Println("TWO!")
		// } else if v%3 == 0 {
		// 	fmt.Println("THREE!")
		// } else if v%2 == 0 && v%3 == 0 {
		// 	fmt.Println("SIX!")
		// } else {
		// 	fmt.Println("nevermind...")
		// }
	}

}

// SHADOWING ISSUE
//
//	func exercise3() {
//		var total int
//		for i := 0; i < 10; i++ {
//			total := total + i
//			fmt.Println(total)
//		}
//		fmt.Println(total)
//	}

// CORRECT
func exercise3() {
	var total int
	for i := 0; i < 10; i++ {
		total = total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}
