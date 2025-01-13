package main

import "fmt"

func main() {

	var max = 3
	var numbers = []int{1, 2, 3, 4, 5, 3, 2, 3, 4, 5, 3, 2, 5, 5}
	var bnyak, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("found \t:", bnyak)
	fmt.Println("value \t:", theNumbers)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
