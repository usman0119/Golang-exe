package main

import (
	"fmt"
	"maps"
)

func main() {

	n := map[string]int{"foo": 1, "bar": 3}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	} else {
		fmt.Println("bukan")
	}
}
