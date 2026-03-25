package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	for i := 0; i < len(pow); i++ {
		fmt.Printf("2**%d = %d\n", i, pow[i])
	}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
