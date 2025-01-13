package main

import "fmt"

func main() {
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			fmt.Println("berikut bilangan genap = ", j)
		} else if j%1 == 0 {
			fmt.Println(" ganjil = ", j)
		}
	}
}
