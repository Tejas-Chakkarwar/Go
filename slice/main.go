package main

import (
	"fmt"
	"strings"
)

func main() {
	primes := [6]int{2, 3, 4, 5, 6, 7}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"Tejas",
		"Anuja",
		"Keyur",
		"Atharva",
	}

	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{3, 4, 5, 6, 7}
	fmt.Println(q)

	r := []bool{true, false, false, true}
	fmt.Println(r)

	e := []struct {
		j int
		d bool
	}{
		{2, true},
		{3, false},
		{4, false},
		{5, true},
	}

	fmt.Println(e)

	z1 := make([]int, 5)
	printSlice("z1", z1)

	z2 := make([]int, 0, 5)
	printSlice("z2", z2)

	z3 := z2[:2]
	printSlice("z3", z3)

	z4 := z3[2:5]
	printSlice("z4", z4)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[1][1] = "X"
	board[2][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
