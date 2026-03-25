package main

import "fmt"

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

}
