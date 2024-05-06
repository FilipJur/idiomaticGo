package main

func main() {
	//––––––––––––––––––––––––––
	// !!!!!1. Shadowing!!!!!
	//––––––––––––––––––––––––––
	// x := 10
	// if x > 5 {
	// 	x, y := 5, 20
	// 	fmt.Println(x, y)
	// }
	// fmt.Println(x)
	//––––––––––––––––––––––––––
	// !!!!!2. If statements!!!!!
	//––––––––––––––––––––––––––
	// if n := rand.Intn(10); n == 0 {
	// 	fmt.Println("Thats too low")
	// } else if n > 5 {
	// 	fmt.Println("Thats too big:", n)
	// } else {
	// 	fmt.Println("Thats a good number:", n)
	// }
	//––––––––––––––––––––––––––
	// !!!!!3. For - 4 ways !!!!!!!!!!!!
	//––––––––––––––––––––––––––
	// 1)Complete for statement
	//––––––––––––––––––––––––––
	// for i := 0; i < 10; {
	// 	fmt.Println(i)
	// 	if i%2 == 0 {
	// 		i++
	// 	} else {
	// 		i += 2
	// 	}
	// }
	//––––––––––––––––––––––––––
	// 2)Condition only for statement
	//––––––––––––––––––––––––––
	// i := 1
	// for i < 100 {
	// 	i = i * 2
	// 	fmt.Println(i)
	// }
	//––––––––––––––––––––––––––
	// 3)Infinite for statement
	//––––––––––––––––––––––––––
	// for {
	// 	fmt.Println("HELLO KOKOTE")
	// 	break
	// }
	// i := 0
	// for {
	// 	i++
	// 	fmt.Println("ČAU TY KOKOTE! :D", i)
	// 	if !(i < 10) {
	// 		break
	// 	}
	// }
	// ----- WITHOUT CONTINUE -------
	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		if i%5 == 0 {
	// 			fmt.Println("FizzBuzz")
	// 		} else {
	// 			fmt.Println("Fizz")
	// 		}
	// 	} else if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }
	// ------- WITH CONTINUE ------------
	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("Fizzbuzz")
	// 		continue
	// 	}
	// 	if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 		continue
	// 	}
	// 	if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//––––––––––––––––––––––––––
	// 4)For range statement
	//––––––––––––––––––––––––––
	// evenVals := []int{2, 4, 5, 8, 10, 12}
	// for i, v := range evenVals {
	// 	fmt.Println(i, v)
	// }
	// evenVals := []int{2, 4, 5, 8, 10, 12}
	// for _, v := range evenVals {
	// 	fmt.Println(v)
	// }
	// uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	// for k := range uniqueNames {
	// 	fmt.Println(k)
	// }
	// m := map[string]int{
	// 	"a": 1,
	// 	"c": 3,
	// 	"b": 2,
	// }

	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Loop", i)
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }

	// OVER STRINGS
	// samples := []string{"hello", "apple_n!"}
	// for _, sample := range samples {
	// 	for i, r := range sample {
	// 		fmt.Println(i, r, string(r))
	// 	}
	// }
	// fmt.Println()
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	v *= 2
	// }
	// fmt.Println(evenVals)

	// -------Labeling Your for Statements----------
	// 	samples := []string{"Hello", "apple_n!"}
	// outer:
	// 	for _, sample := range samples {
	// 		for i, r := range sample {
	// 			fmt.Println(i, r, string(r))
	// 			if r == 'l' {
	// 				continue outer
	// 			}
	// 		}
	// 	}
	// 	fmt.Println()

	// evenVals := []int{2, 4, 6, 8, 10}
	// for i, v := range evenVals {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	if i == len(evenVals)-1 {
	// 		break
	// 	}
	// 	fmt.Println(i, v)
	// }

	// evenVals := []int{2, 4, 6, 8, 10}
	// for i := 1; i < len(evenVals)-1; i++ {
	// 	fmt.Println(i, evenVals[i])
	// }

	// ------ SWITCH ---------
	// words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	// for _, word := range words {
	// 	switch size := len(word); size {
	// 	case 1, 2, 3, 4:
	// 		fmt.Println(word, "is a short word!")
	// 	case 5:
	// 		wordLen := len(word)
	// 		fmt.Println(word, "is exactly the right lenght: ", wordLen)
	// 	case 6, 7, 8, 9:
	// 	default:
	// 		fmt.Println(word, "is a long word!")
	// 	}
	// }

	// ------ BREAK OUT OF SWITCH ---------
	// loop:
	// 	for i := 0; i < 10; i++ {
	// 		switch i {
	// 		case 0, 2, 4, 6:
	// 			fmt.Println(i, "is even")
	// 		case 3:
	// 			fmt.Println(i, "is divisible by 3 but not by 2")
	// 		case 7:
	// 			fmt.Println("Exit the loop!")
	// 			break loop
	// 		default:
	// 			fmt.Println(i, "is boring")
	// 		}
	// 	}

	// ------ BLANK SWITCHES ---------
	// words := []string{"hi", "salutatuions", "hello"}
	// for _, word := range words {
	// 	switch wordLen := len(word); {
	// 	case wordLen < 5:
	// 		fmt.Println(word, "is a short word")
	// 	case wordLen > 10:
	// 		fmt.Println(word, "is a long word")
	// 	default:
	// 		fmt.Println(word, "is exactly the right length")
	// 	}
	// }

	// -------------Choosing Between if and switch --------------
	// for i := 1; i < 100; i++ {
	// 	switch {
	// 	case i%3 == 0 && i%5 == 0:
	// 		fmt.Println("Fizzbuz")
	// 	case i%3 == 0:
	// 		fmt.Println("Fizz")
	// 	case i%5 == 0:
	// 		fmt.Println("Buzz")
	// 	default:
	// 		fmt.Println(i)
	// 	}
	// }

	// ------------- GOTO --------------
	// EXAMPLE - invalid code:
	// 	a := 10
	// 	b := 20
	// skip:
	// 	c := 30
	// 	fmt.Println(a, b, c)
	// 	if c > a {
	// 		goto inner
	// 	}
	// 	if a < b {
	// 	inner:
	// 		fmt.Println("a is less than b")
	// 	}
	// EXAMPLE - valid code:
	// 	a := rand.Intn(10)
	// 	for a < 100 {
	// 		if a%5 == 0 {
	// 			goto done
	// 		}
	// 		a = a*2 + 1
	// 	}
	// 	fmt.Println("Do something when the loop completes normally")
	// done:
	// 	fmt.Println("Do complicated stuff no matter why we left the loop")
	// 	fmt.Println(a)

	// EXERCISES!!!!!
	exercisesCh4()
}
