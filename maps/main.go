package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68, -74.39,
	}
	fmt.Println(m["Bell Labs"])

	z := make(map[string]int)

	z["Answer"] = 42
	fmt.Println("The value:", z["Answer"])

	z["Answer"] = 48
	fmt.Println("The value:", z["Answer"])

	for key, value := range z {
		fmt.Println(key, value)
	}

	delete(z, "Answer")
	fmt.Println("The value:", z["Answer"])

	v, ok := z["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
