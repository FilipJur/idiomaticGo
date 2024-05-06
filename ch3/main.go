package main

func main() {

	/*
	   	x := []int{1, 2, 3, 4, 5}
	   	y := []int{1, 2, 3, 4, 5}
	   	z := []int{1, 2, 3, 4, 5, 6}

	   	fmt.Println(slices.Equal(x, y))
	   	fmt.Println(slices.Equal(x, z))

	   	var x = []int{1, 2, 3}
	   	var y = []int{4, 5, 6}
	   	x = append(x, y...)
	   	fmt.Println(cap(x))
	   	var x []int
	   	fmt.Println(x, len(x), cap(x))

	   	x = append(x, 10)
	   	fmt.Println(x, len(x), cap(x))

	   	x = append(x, 20)
	   	fmt.Println(x, len(x), cap(x))

	   	x = append(x, 30)
	   	fmt.Println(x, len(x), cap(x))

	   	x = append(x, 40)
	   	fmt.Println(x, len(x), cap(x))

	   	x = append(x, 50)
	   	fmt.Println(x, len(x), cap(x))

	   	x = append(x, 60)
	   	fmt.Println(x, len(x), cap(x))

	   	x := make([]int, 0, 15)

	   	x = append(x, 10)
	   	fmt.Println(x, len(x), cap(x))

	   	s := []string{"First", "Second", "Third"}
	   fmt.Println(s, len(s))
	   clear(s)
	   fmt.Println(s, len(s))
	*/

	// x := []int{1, 2, 3, 4}
	// y := make([]int, 2)
	// num := copy(y, x[2:])
	// fmt.Println(y, num)

	// x := []int{1, 2, 3, 4}
	// d := [4]int{5, 6, 7, 8}
	// y := make([]int, 2)
	// copy(y, d[:])
	// fmt.Println(y)
	// copy(d[:], x)
	// fmt.Println(d)

	// xSlice := []int{1, 2, 3, 4}
	// xArray := [4]int(xSlice)
	// smallArray := [2]int(xSlice)
	// xSlice[0] = 10
	// fmt.Println(xSlice, xArray, smallArray)

	// var s string = "Hello There"
	// var b byte = s[6]
	// fmt.Println(b)

	// nilMap := map[string][]string{
	// 	"Orcas":   []string{"Fred", "Ralph", "Bijous"},
	// 	"Lion":    []string{"Sarah", "Peter", "Billie"},
	// 	"Kittens": []string{"Waldo", "Raul", "Ze"},
	// }

	// ages := make(map[int][]string, 10)

	// fmt.Println(nilMap, len(ages), ages)

	// totalWins := map[string]int{}
	// totalWins["Orcas"] = 1
	// totalWins["Lions"] = 2

	// fmt.Println(totalWins["Orcas"])
	// fmt.Println(totalWins["Kittens"])

	// totalWins["Kittens"]++
	// fmt.Println(totalWins["Kittens"])

	// totalWins["Lions"] = 3
	// fmt.Println(totalWins[])

	// m := map[string]string{
	// 	"hello": "5",
	// 	"world": "0",
	// }
	// v, ok := m["hello"]
	// fmt.Println(v, ok)

	// v, ok = m["world"]
	// fmt.Println(v, ok)

	// v, ok = m["goodbye"]
	// fmt.Println(v, ok)

	// n := map[string]int{
	// 	"hello": 5,
	// 	"world": 10,
	// }

	// clear(n)
	// fmt.Println(n)

	// intSet := map[int]bool{}
	// vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	// for _, v := range vals {
	// 	intSet[v] = true
	// }
	// fmt.Println(len(vals), len(intSet))
	// fmt.Println(intSet[5])
	// fmt.Println(intSet[500])
	// if intSet[10] {
	// 	fmt.Println("100 is in the set")
	// }

	type person struct {
		name string
		age  string
		pet  string
	}
}
